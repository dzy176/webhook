package main

import (
	"log"
	"net/http"
	"os/exec"

	"go.uber.org/zap"
)

var logger *zap.Logger

func push(w http.ResponseWriter, r *http.Request) {
	logger.Info("sync files from git")
	cmd := exec.Command("/bin/sh", "update.sh")
	err := cmd.Run()
	if err != nil {
		logger.Error(err.Error())
	}
}

func main() {
	// 设置路由，如果访问/，则调用index方法
	logger, _ = zap.NewDevelopment()
	logger.Info("webhook start at 47.101.47.118:8080")
	http.HandleFunc("/push", push)

	// 启动web服务，监听9090端口
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
