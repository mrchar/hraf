package hraf

import (
	"context"
	"errors"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

type key string

const (
// keyRequest        key = "keyRequest"
// keyResponseWriter key = "keyResponseWriter"
// keyEnabledDebug   key = "keyEnabledDebug"
// keyEncoder        key = "keyEncoder"
)

// Context Context包含 http.Request http.ResponseWriter
type Context struct {
	context.Context
	server         *Server
	request        *http.Request
	responseWriter http.ResponseWriter
	encoder        Encoder
}

// GetRequest 返回http.Request
func (c *Context) GetRequest() *http.Request {
	return c.request
}

// GetResponseWriter 返回http.ResponseWriter
func (c *Context) GetResponseWriter() http.ResponseWriter {
	return c.responseWriter
}

// Params 从http请求中获取参数并填入v中
func (c *Context) Params(v interface{}) error {
	// 读取请求
	request := c.GetRequest()

	if c.encoder == nil {
		err := errors.New("The server doesn't have an encoder")
		log.Panic(err)
	}

	// 解码请求
	err := c.encoder.Decode(request, v)
	if err != nil {
		if c.Debug() {
			log.Println(err)
		}
		return err
	}
	return nil
}

// Respond 解析内容到 Respond 中
func (c *Context) Respond(v interface{}) error {
	return ErrNotImplement
}

// File 使用一个文件作为相应
// TODO: http cache
// TODO: set header
func (c *Context) File(name string) error {
	file, err := os.Open(name)
	if err != nil {
		if c.server.Debug() {
			log.Println(err)
		}
		return err
	}

	defer file.Close()

	_, err = io.Copy(c.responseWriter, file)
	if err != nil {
		if c.server.Debug() {
			log.Println(err)
		}
		return err
	}

	return nil
}

// Error 想Response中写入一个错误并返回
// XXX 使用encoder编码错误
// TODO TraceStack
func (c *Context) Error(code int, err error) {
	writer := c.GetResponseWriter()
	var msg string
	if c.Debug() {
		for err != nil {
			msg = msg + err.Error()
			if !strings.HasSuffix(msg, "\n") {
				msg = msg + "\n"
			}
			err = errors.Unwrap(err)
		}
	} else {
		msg = err.Error()
	}

	writer.WriteHeader(code)
	writer.Write([]byte(msg))
}

// Debug 是否开启了Debug模式
func (c *Context) Debug() bool {
	return c.server.Debug()
}
