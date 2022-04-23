package main

import (
	"github.com/cncamp-tasks/miniserver/internal/config"
	"github.com/cncamp-tasks/miniserver/internal/handler"
	"github.com/cncamp-tasks/miniserver/internal/log"
	"github.com/sirupsen/logrus"
	"net/http"
)

func initApp() error {
	// 配置初始化
	// 配置加载

	// 日志初始化
	log.InitLog()
	return nil
}

func main() {
	err := initApp()
	if err != nil {
		return
	}

	// 运行服务器
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler.RootHandler)

	logrus.Infof("miniserver start listen[%s]...", config.AppBindAddress)
	if err := http.ListenAndServe(config.AppBindAddress, mux); err != nil {
		logrus.Fatalf("miniserver listen err: %s", err)
	}
}
