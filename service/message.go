package service

import (
	"context"
	"encoding/json"
	"encoding/xml"
	"github.com/baidubce/bce-qianfan-sdk/go/qianfan"
	"io"
	"log"
	"net/http"
	"time"
)

// MessageHandler
func MessageHandler(w http.ResponseWriter, r *http.Request) {

	bs, _ := io.ReadAll(r.Body)
	msg := NewMsg(bs)
	if msg == nil {
		log.Printf("log body is nil")
		return
	}

	log.Printf("log body : %s", msg.Content)

	// 发起对话，例如介绍下北京
	resp, err := chatClient.Do(
		context.TODO(),
		&qianfan.ChatCompletionRequest{
			Messages: []qianfan.ChatCompletionMessage{
				qianfan.ChatCompletionUserMessage(msg.Content),
			},
		},
	)
	if err != nil {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.Write(msg.GenerateEchoData(err.Error()))
		return
	}

	log.Printf("qianfan body : %s", resp.Result)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write(msg.GenerateEchoData(resp.Result))
}

type Msg struct {
	ToUserName   string `xml:"ToUserName"`
	FromUserName string `xml:"FromUserName"`
	CreateTime   int64  `xml:"CreateTime"`
	MsgType      string `xml:"MsgType"`
	Event        string `xml:"Event"`
	Content      string `xml:"Content"`
	Recognition  string `xml:"Recognition"`

	MsgId int64 `xml:"MsgId,omitempty"`
}

func NewMsg(data []byte) *Msg {
	var msg Msg
	if err := json.Unmarshal(data, &msg); err != nil {
		return nil
	}
	return &msg
}

func (msg *Msg) GenerateEchoData(s string) []byte {
	data := Msg{
		ToUserName:   msg.FromUserName,
		FromUserName: msg.ToUserName,
		CreateTime:   time.Now().Unix(),
		MsgType:      "text",
		Content:      s,
	}
	bs, _ := xml.Marshal(&data)
	return bs
}
