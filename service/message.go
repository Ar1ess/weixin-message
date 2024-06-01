package service

import (
	"encoding/json"
	"fmt"
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
		if err != nil {
			fmt.Fprint(w, "内部错误")
		}
		w.Write(msg)
		return
	}

	fmt.Printf("fmt body : %s", string(body))
	log.Printf("log body : %s", string(body))

	res.Code = 0
	res.Data = ""
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
