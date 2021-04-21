package definition

import "net/http"

type Codec interface {
	// 将golang中的数据编码为字节数组
	Encode(interface{}) ([]byte, error)
	// 将字节数组解码到golang的变量中
	Decode([]byte, interface{}) error
	Pluck(*http.Request, interface{}) error
	Respond(http.ResponseWriter, interface{}) error
}
