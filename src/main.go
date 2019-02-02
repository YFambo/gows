package main

import (
	"net/http"
	"impl"
	_ "net/http/pprof"
)

func main() {
	hub := impl.NewHub()
	go hub.Run()
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		impl.Handler(hub, w, r)
	})
	impl.SystemOutput.Info("服务开启成功！")
	http.ListenAndServe(":7777", nil)
}
