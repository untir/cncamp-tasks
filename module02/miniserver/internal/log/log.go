package log

import (
	"github.com/cncamp-tasks/miniserver/internal/config"
	"github.com/sirupsen/logrus"
)

// InitLog 初始化日志
func InitLog() {
	logrus.SetLevel(config.AppLogLevel)
	logrus.SetReportCaller(true)
	logrus.SetFormatter(&logrus.JSONFormatter{})
}
