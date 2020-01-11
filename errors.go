package hraf

import (
	"errors"
)

var (
	// ErrNotImplement 一个函数返回ErrNotImplement
	// 表示这个函数的功能还没有实现
	ErrNotImplement = errors.New("not implement")
)
