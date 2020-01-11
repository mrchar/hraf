package hraf

import "net/http"

// Encoder 用来编码响应，解码请求
type Encoder interface {
	// Decode 将http.Request中的内容解析到interface中
	Decode(r *http.Request, v interface{}) error
	// Encode 将要返回的结果写入Response
	Encode(w http.ResponseWriter, v interface{}) error
}
