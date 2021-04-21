package router

import "net/http"

type Router struct {
	*http.ServeMux
}

func New() Router {
	return Router{
		new(http.ServeMux),
	}
}
