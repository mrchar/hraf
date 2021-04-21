package server

import (
	"context"
	"net/http"

	def "github.com/mrchar/hraf/definition"
)

// Server Server实现一个http服务器，用于接收请求、处理、并返回结果
type Server struct {
	address string
	server  http.Server
	router  def.Router
}

// New New 构建一个Server
func New(address string, router def.Router) *Server {
	return &Server{
		address: address,
		router:  router,
	}
}

func (s *Server) Start() error {
	if s.address != "" {
		s.server.Addr = s.address
	}

	s.server.Handler = s.router

	return s.server.ListenAndServe()
}

func (s *Server) Stop(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}

func (s *Server) HandleFunc(pattern string, f func(http.ResponseWriter, *http.Request)) {
	s.router.HandleFunc(pattern, f)
}
