package handler

import (
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"strings"
)

const (
	VersionEnvKey = "VERSION"
)

// RootHandler 根路径请求处理函数
func RootHandler(w http.ResponseWriter, r *http.Request) {
	statusCode := http.StatusOK
	defer func() {
		w.WriteHeader(statusCode)
		accessLog(r, statusCode)
	}()

	// 将 request header 写入 response header
	for k, v := range r.Header {
		for _, vv := range v {
			w.Header().Set(k, vv)
		}
	}

	// 读取VERSION
	version := os.Getenv(VersionEnvKey)
	w.Header().Set(VersionEnvKey, version)

	// 返回响应数据
	_, _ = w.Write([]byte("root ok"))
}

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	statusCode := http.StatusOK
	defer func() {
		w.WriteHeader(statusCode)
		accessLog(r, statusCode)
	}()

	// 将 request header 写入 response header
	for k, v := range r.Header {
		for _, vv := range v {
			w.Header().Set(k, vv)
		}
	}

	// 返回响应数据
	_, _ = w.Write([]byte("health ok"))
}

func accessLog(r *http.Request, statusCode int) {
	logrus.WithFields(logrus.Fields{
		"ip":      getSrcIP(r),
		"request": r.RequestURI,
		"status":  statusCode,
	}).Info()
}

func getSrcIP(r *http.Request) string {
	ip := r.Header.Get("X-Real-IP")
	if ip == "" {
		ips := strings.Split(r.RemoteAddr, ":")
		if len(ips) > 0 {
			ip = ips[0]
		}
	}
	return ip
}
