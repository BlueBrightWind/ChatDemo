package socket

import (
	"ChatDemo/sql"
	"encoding/json"
	"time"
)

type MessageManager struct {
	MessageByteChan chan []byte
}

func NewMessageManager() *MessageManager {
	manager := &MessageManager{
		MessageByteChan: make(chan []byte),
	}
	go manager.AutoDispatchMessage()
	return manager
}

func (manager *MessageManager) AddMessage(msg []byte) {
	manager.MessageByteChan <- msg
}

func (manager *MessageManager) AutoDispatchMessage() {
	for {
		select {
		case msg := <-manager.MessageByteChan:
			parsedMsg, err := ParseMessage(msg)
			if err != nil {
				continue
			}
			parsedMsg.CreateTime = uint64(time.Now().Unix())
			// Use parsedMsg here
			if parsedMsg.Type == 1 {
				// Private message
				manager.DispatchPrivateMessage(parsedMsg)
			} else if parsedMsg.Type == 2 {
				manager.DispatchGroupMessage(parsedMsg)
				// Group message
			}
		}
	}
}

func (manager *MessageManager) DispatchPrivateMessage(parsedMsg *Message) {
	msg, err := json.Marshal(parsedMsg)
	if err != nil {
		return
	}
	NodeManagerApi.SetMessage(parsedMsg.TargetId, msg)
}

func (manager *MessageManager) DispatchGroupMessage(parsedMsg *Message) {
	usersId := sql.GetUsersIdByCommunityId(parsedMsg.TargetId)
	msg, err := json.Marshal(parsedMsg)
	if err != nil {
		return
	}
	for _, userId := range usersId {
		NodeManagerApi.SetMessage(userId, msg)
	}
}

var MessageManagerApi = NewMessageManager()
