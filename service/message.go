package service

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/baidubce/bce-qianfan-sdk/go/qianfan"
	"io"
	"log"
	"net/http"
)

// MessageHandler
func MessageHandler(w http.ResponseWriter, r *http.Request) {
	res := &JsonResult{}

	w.Header().Set("content-type", "application/json")

	body, err := PeekRequest(r)
	if err != nil {
		res.Code = -1
		res.ErrorMsg = err.Error()
		msg, err := json.Marshal(res)
		log.Printf("err: %s", err.Error())
		if err != nil {
			fmt.Fprint(w, "内部错误")
		}
		w.Write(msg)
		return
	}

	m := &Model{}
	if err := json.Unmarshal(body, &m); err != nil {
		res.Code = -1
		res.ErrorMsg = err.Error()
		msg, err := json.Marshal(res)
		log.Printf("err: %s", err.Error())
		if err != nil {
			fmt.Fprint(w, "内部错误")
		}
		w.Write(msg)
		return
	}
	log.Printf("log body : %s", string(body))

	// 发起对话，例如介绍下北京
	resp, err := chatClient.Do(
		context.TODO(),
		&qianfan.ChatCompletionRequest{
			Messages: []qianfan.ChatCompletionMessage{
				qianfan.ChatCompletionUserMessage(m.Content),
			},
		},
	)
	if err != nil {
		res.Code = -1
		res.ErrorMsg = err.Error()
		log.Printf("err: %s", err.Error())
		msg, err := json.Marshal(res)
		if err != nil {
			fmt.Fprint(w, "内部错误")
		}
		w.Write(msg)
		return
	}

	res.Code = 0
	res.Data = resp.Result
	res.ErrorMsg = ""
	msg, err := json.Marshal(res)
	if err != nil {
		fmt.Fprint(w, "内部错误")
		return
	}
	w.Header().Set("content-type", "application/json")
	w.Write(msg)
}

func PeekRequest(request *http.Request) ([]byte, error) {
	if request.Body != nil {
		byts, err := io.ReadAll(request.Body) // io.ReadAll as Go 1.16, below please use ioutil.ReadAll
		if err != nil {
			return nil, err
		}
		return byts, nil
	}

	return make([]byte, 0), nil
}

type Model struct {
	MsgId        int64  `json:"MsgId"`
	Content      string `json:"Content"`
	MsgType      string `json:"MsgType"`
	CreateTime   int64  `json:"CreateTime"`
	FromUserName string `json:"FromUserName"`
	ToUserName   string `json:"ToUserName"`
}
