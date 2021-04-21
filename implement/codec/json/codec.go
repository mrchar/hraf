package json

import "encoding/json"

var def Codec

type Codec struct{}

func (Codec) Encode(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

func (Codec) Decode(bytes []byte, v interface{}) error {
	return json.Unmarshal(bytes, v)
}

func Encode(v interface{}) ([]byte, error) {
	return def.Encode(v)
}

func Decode(bytes []byte, v interface{}) error {
	return def.Decode(bytes, v)
}
