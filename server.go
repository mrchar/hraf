package hraf

import (
	"fmt"
	"log"
	"net/http"

	"github.com/mrchar/hraf/encoding/json"
)

var stdEncoder = &json.Encoder{}

// Server Server实现一个http服务器，用于接收请求、处理、并返回结果
type Server struct {
	http.Server
	encoder     Encoder
	controllers map[string]Controller
	debug       bool
}

// NewServer NewServer 构建一个Server
func NewServer() *Server {
	return &Server{
		encoder: stdEncoder,
	}
}

// // Party 创建分组
// func (s *Server) Party() {

// }

// Register 注册controllers到Sever上
func (s *Server) Register(route string, controller Controller) {
	if _, ok := s.controllers[route]; ok {
		err := fmt.Errorf("There is an controller with the same name")
		log.Fatal(err)
	}
	s.controllers[route] = controller
}

// func (s *Server)Route()  {

// }

// File 添加一个文件作为 Get route 的响应
// func (s *Server) File(route string, name string) {

// }

// Build 处理注册的Controller并创建http.Handler
func (s *Server) Build() (http.Handler, error) {
	return &Route{}, ErrNotImplement
}

// initialize 初始化服务器
// 检查Controller
// 创建路由
func (s *Server) initialize(addr string, debug ...bool) error {
	s.Server = http.Server{
		Addr: addr,
	}

	handler, err := s.Build()
	if err != nil {
		if s.debug {
			log.Println(err)
		}
		return err
	}
	s.Server.Handler = handler

	s.debug = EnvDebugEnable
	if len(debug) > 0 {
		s.debug = debug[0]
	}

	return ErrNotImplement
}

// Run 启动服务器
func (s *Server) Run(addr string, debug ...bool) error {
	err := s.initialize(addr, debug...)
	if err != nil {
		err := fmt.Errorf("Failed to initialize the server with err: %w ", err)
		log.Fatal(err)
	}
	return s.ListenAndServe()
}

// Stop 停止服务器
// func (s *Server) Stop() error {
// 	return s.Close()
// }

// SetEncoder 设置编码器
// 编码器用于将返回的内容编码到 http response 中
// 以及将 http request 中的内容解码到结构体中
func (s *Server) SetEncoder(encoder Encoder) {
	s.encoder = encoder
}

// EnableDebug 是否开启调试模式
func (s *Server) EnableDebug(debug bool) {
	s.debug = debug
}

// Debug 获取Server是否开启Debug模式
func (s *Server) Debug() bool {
	return s.Debug()
}
