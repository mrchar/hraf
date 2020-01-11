package hraf

import "os"

var (
	// EnvDebugEnable 是否开启调试模式
	EnvDebugEnable = false
)

func init() {
	if os.Getenv("HRAF_DEBUG") == "true" {
		EnvDebugEnable = true
	}
}
