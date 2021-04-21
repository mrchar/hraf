package definition

import (
	"context"
	"net/http"
)

// Server HTTP RESTFUL API Server
type Server interface {
	// 启动服务器
	Start() error
	// 优雅关闭服务器
	Stop(context.Context) error
	// 当路由匹配路径时，使用给出的函数处理请求
	HandleFunc(pattern string, f func(http.ResponseWriter, *http.Request))
	Handle(pattern string, handler http.Handler)
}
