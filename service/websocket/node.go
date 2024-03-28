package socket

import (
	"ChatDemo/global"
	"log"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

type Node struct {
	ID                    uint            //用户ID
	Conn                  *websocket.Conn //连接
	HeartbeatTimer        *time.Timer
	HeartbeatChannel      chan bool
	DataQueue             chan []byte
	Exist                 bool
	PublishProcessorQueue chan uint
	Locker                sync.RWMutex
}

func NewNode(id uint, conn *websocket.Conn, manageQueue chan uint) *Node {
	node := &Node{
		ID:                    id,
		Conn:                  conn,
		HeartbeatTimer:        nil,
		HeartbeatChannel:      nil,
		DataQueue:             nil,
		Exist:                 true,
		PublishProcessorQueue: manageQueue,
		Locker:                sync.RWMutex{},
	}
	node.ID = id

	heartbeatTimeout := time.Duration(global.Config.SocketConfig.HeartbeatMaxTime) * time.Second
	node.HeartbeatTimer = time.NewTimer(heartbeatTimeout)

	node.HeartbeatChannel = make(chan bool)
	node.DataQueue = make(chan []byte)
	node.Exist = true
	node.PublishProcessorQueue = manageQueue

	conn.SetPongHandler(func(appData string) error {
		node.HeartbeatChannel <- true
		return nil
	})

	node.Conn = conn

	go node.AutoHeartBeatSend()
	go node.AutoHeartBeatReceive()
	go node.AutoProcessData()
	go node.AutoReceiveData()

	return node
}

func (node *Node) DeteleNode() {
	node.Locker.Lock()
	if !node.Exist {
		node.Locker.Unlock()
		return
	}
	log.Println("节点关闭, ID:", node.ID)
	node.Exist = false
	node.Conn.Close()
	node.DataQueue <- []byte("") // 向队列写入数据，防止数据处理协程阻塞无法关闭
	node.PublishProcessorQueue <- node.ID
	// close(node.HeartbeatChannel)
	// close(node.DataQueue)
	node.Locker.Unlock()
}

func (node *Node) AutoHeartBeatSend() {
	// defer log.Println("心跳发送协程关闭")

	defer node.DeteleNode()

	heartbeatInterval := time.Duration(global.Config.SocketConfig.HeartBeatInterval) * time.Second
	heartbeatTimeOut := time.Duration(global.Config.SocketConfig.HeartbeatMaxTime) * time.Second

	for {
		err := node.Conn.WriteControl(websocket.PingMessage, nil, time.Now().Add(heartbeatTimeOut))
		if err != nil {
			return
		}
		time.Sleep(heartbeatInterval)
	}
}

func (node *Node) AutoHeartBeatReceive() {
	// defer log.Println("心跳检测协程关闭")

	defer node.DeteleNode()

	defer node.HeartbeatTimer.Stop()

	heartbeatTimeout := time.Duration(global.Config.SocketConfig.HeartbeatMaxTime) * time.Second

	node.HeartbeatTimer.Reset(heartbeatTimeout)

	for {
		select {
		case <-node.HeartbeatTimer.C:
			return
		case <-node.HeartbeatChannel:
			if !node.HeartbeatTimer.Stop() {
				<-node.HeartbeatTimer.C
			}
			node.HeartbeatTimer.Reset(heartbeatTimeout)
		}
	}
}

func (node *Node) AutoProcessData() {
	// defer log.Println("数据处理协程关闭")
	defer node.DeteleNode()
	for {
		select {
		case data := <-node.DataQueue:
			err := node.Conn.WriteMessage(websocket.TextMessage, data)
			if err != nil {
				return
			}
		}
	}
}

func (node *Node) AutoReceiveData() {
	// defer log.Println("消息接收协程关闭")
	defer node.DeteleNode()
	for {
		_, p, err := node.Conn.ReadMessage()
		if err != nil {
			return
		}
		MessageManagerApi.AddMessage(p)
	}
}
