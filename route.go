package hraf

import "net/http"

// Route 实现一个http.Handler
type Route struct{}

func (r *Route) ServeHTTP(writer http.ResponseWriter, request *http.Request) {

}
