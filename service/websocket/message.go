package socket

import "encoding/json"

type Message struct {
	UserId     uint   //发送者
	TargetId   uint   //接受者
	Type       int    //发送类型  1私聊  2群聊
	Media      int    //消息类型  1文字 2表情包 3语音 4图片 /表情包
	Content    string //消息内容
	CreateTime uint64 //创建时间
}

func ParseMessage(data []byte) (*Message, error) {
	message := Message{}
	err := json.Unmarshal(data, &message)
	if err != nil {
		return nil, err
	}
	return &message, nil
}
