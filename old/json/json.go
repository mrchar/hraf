package json

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"reflect"
	"strings"
)

// 将字典编码到响应中时，使用这些键获取要编码到响应中的内容
const (
	Header  = "header"
	Code    = "code"
	Status  = "status"
	Message = "message"
	Data    = "data"
)

var (
	// ErrNotImplement 一个函数返回ErrNotImplement
	// 表示这个函数的功能还没有实现
	ErrNotImplement = errors.New("not implement")
	// ErrInvalidType 类型错误
	ErrInvalidType = errors.New("invalid type")
)

// Encoder 用于从http请求中解码内容，已经编码内容到http响应中
type Encoder struct{}

// Decode 从http请求中解码内容
func (e *Encoder) Decode(r *http.Request, v interface{}) error {
	tv := reflect.TypeOf(v)
	// v 必须是指针类型
	if tv.Kind() != reflect.Ptr {
		return ErrInvalidType
	}

	iv := reflect.Indirect(reflect.ValueOf(v))

	switch iv.Kind() {
	case reflect.Struct:
		err := e.decodeStruct(r, iv.Interface())
		return err
	default:
		return ErrInvalidType
	}
}

func (e *Encoder) decodeStruct(r *http.Request, v interface{}) error {
	return ErrNotImplement
}

// Encode 将内容编码到http响应中
func (e *Encoder) Encode(w http.ResponseWriter, v interface{}) error {
	// 如果v是io.Reader, 则直接将内容写入正文
	if reader, ok := v.(io.Reader); ok {
		_, err := io.Copy(w, reader)
		if err != nil {
			err := fmt.Errorf("An error occurred while copying the content: %w", err)
			return err
		}
		return nil
	}

	tv := reflect.TypeOf(v)
	switch tv.Kind() {
	case reflect.Ptr:
		val := reflect.ValueOf(v)
		err := e.Encode(w, reflect.Indirect(val).Interface())
		if err != nil {
			err = fmt.Errorf("An error occurred while encoding the content into the response: %w", err)
			return err
		}
	case reflect.Struct:
		return e.encodeStruct(w, v)
	case reflect.Map:
		m, ok := v.(map[string]interface{})
		if !ok {
			err := fmt.Errorf("only can use map[string]interface{}, error: %w", ErrInvalidType)
			return err
		}
		err := e.encodeMap(w, m)
		if err != nil {
			err = fmt.Errorf("An error occurred while encoding map into the response: %w", err)
			return err
		}
	default:
		return ErrInvalidType
	}
	return nil
}

func (e *Encoder) encodeStruct(w http.ResponseWriter, v interface{}) error {
	val := reflect.ValueOf(v)
	n := val.NumField()
	for i := 0; i < n; i++ {
		// field := val.Field(i)
	}
	return ErrNotImplement
}

func (e *Encoder) encodeMap(w http.ResponseWriter, v map[string]interface{}) error {
	header, ok := v[Header]
	if ok {
		kvs, ok := header.(map[string]interface{})
		if !ok {
			err := fmt.Errorf("Unable to encode %v into header", header)
			return err
		}
		for key, val := range kvs {
			tv := reflect.TypeOf(val)
			switch tv.Kind() {
			case reflect.String:
				w.Header().Add(key, val.(string))
			case reflect.Array:
				vval := reflect.ValueOf(val)
				n := vval.NumField()
				buf := make([]string, n)
				for i := 0; i < n; i++ {
					buf = append(buf, fmt.Sprintf("%v", vval.Field(i).Interface()))
				}
				w.Header().Add(key, strings.Join(buf, "; "))
			default:
				w.Header().Add(key, fmt.Sprintf("%v", val))
			}
		}
	}

	code, ok := v[Code]
	if ok {
		c, ok := code.(int)
		if !ok {
			err := errors.New("code must be int")
			return err
		}
		w.WriteHeader(c)
	}

	// 构建正文内容
	body := struct {
		Status  interface{} `json:"status,omitempty"`
		Message interface{} `json:"message,omitempty"`
		Data    interface{} `json:"data,omitempty"`
	}{
		Status:  v[Status],
		Message: v[Message],
		Data:    v[Data],
	}

	bytes, err := json.Marshal(body)
	if err != nil {
		err := fmt.Errorf("An error occurred while encoding the content: %w", err)
		return err
	}

	_, err = w.Write(bytes)
	if err != nil {
		err := fmt.Errorf("Error while writing content: %w", err)
		return err
	}

	return ErrNotImplement
}
