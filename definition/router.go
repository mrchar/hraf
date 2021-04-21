package definition

import "net/http"

type Router interface {
	// Router 本身也是Handler
	http.Handler
	// 当路由匹配路径时，使用给出的函数处理请求
	HandleFunc(path string, f func(http.ResponseWriter, *http.Request))
}
