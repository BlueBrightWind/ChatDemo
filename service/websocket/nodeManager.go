package socket

import (
	"sync"

	"github.com/gorilla/websocket"
)

type NodeManager struct {
	NodeMap          sync.Map
	NodeRemoveNotice chan uint
}

func NewNodeManager() *NodeManager {
	manager := &NodeManager{
		NodeMap:          sync.Map{},
		NodeRemoveNotice: make(chan uint),
	}

	go manager.AutoRemoveNode()

	return manager
}

func (manager *NodeManager) AddNode(id uint, conn *websocket.Conn) {
	node := NewNode(id, conn, manager.NodeRemoveNotice)
	manager.NodeMap.Store(id, node)
}

func (manager *NodeManager) RemoveNode(id uint) {
	manager.NodeMap.Delete(id)
}

func (manager *NodeManager) AutoRemoveNode() {
	for {
		select {
		case id := <-manager.NodeRemoveNotice:
			manager.RemoveNode(id)
		}
	}
}

func (manager *NodeManager) SetMessage(id uint, message []byte) {
	node, ok := manager.NodeMap.Load(id)
	if !ok {
		return
	}
	node.(*Node).DataQueue <- message
}

var NodeManagerApi = NewNodeManager()
