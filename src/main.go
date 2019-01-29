package main

import (
	"net/http"
	"impl"
)

func main() {
	http.HandleFunc("/ws", impl.Handler)
	impl.SystemOutput.Info("服务开启成功！")
	http.ListenAndServe(":7777", nil)
}
