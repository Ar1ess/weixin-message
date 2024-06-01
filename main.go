package main

import (
	"log"
	"net/http"
	"wxcloudrun-golang/service"
)

func main() {
	//if err := db.Init(); err != nil {
	//	panic(fmt.Sprintf("mysql init failed with %+v", err))
	//}

	service.QianfanInit()

	http.HandleFunc("/", service.IndexHandler)
	http.HandleFunc("/api/count", service.CounterHandler)
	http.HandleFunc("/message", service.MessageHandler)

	log.Fatal(http.ListenAndServe(":80", nil))
}
