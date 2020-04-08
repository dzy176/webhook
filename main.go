package main

import (
	"fmt"
	"log"
	"net/http"
)

func push(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "triger")
}

func main() {
	// 设置路由，如果访问/，则调用index方法
	http.HandleFunc("/push", push)

	// 启动web服务，监听9090端口
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
