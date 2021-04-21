package hraf

import "net/http"

// Controller 用于接收请求、处理、并返回结果
type Controller interface {
	GetRoute() string
	SetRoute(string)
	Handle(Context)
}

// HandleFunc 表示一个可以接收请求、处理、并返回结果的函数
type HandleFunc func(ctx Context)

// Basic 是一个没有实现任何方法的Controller
type Basic struct {
	Route string
}

// GetRoute 获取路径
func (b *Basic) GetRoute() string {
	return b.Route
}

// SetRoute 设置路径
func (b *Basic) SetRoute(route string) {
	b.Route = route
}

// Get Basic.Get
func (b *Basic) Get(ctx Context) {
	ctx.Error(http.StatusMethodNotAllowed, ErrNotImplement)
}

// Head Basic.Head
func (b *Basic) Head(ctx Context) {
	ctx.Error(http.StatusMethodNotAllowed, ErrNotImplement)
}

// Post Basic.Post
func (b *Basic) Post(ctx Context) {
	ctx.Error(http.StatusMethodNotAllowed, ErrNotImplement)
}

// Put Basic.Put
func (b *Basic) Put(ctx Context) {
	ctx.Error(http.StatusMethodNotAllowed, ErrNotImplement)
}

// Patch Basic.Patch
func (b *Basic) Patch(ctx Context) {
	ctx.Error(http.StatusMethodNotAllowed, ErrNotImplement)
}

// Delete Basic.Delete
func (b *Basic) Delete(ctx Context) {
	ctx.Error(http.StatusMethodNotAllowed, ErrNotImplement)
}

// Connect Basic.Connect
func (b *Basic) Connect(ctx Context) {
	ctx.Error(http.StatusMethodNotAllowed, ErrNotImplement)
}

// Options Basic.Options
func (b *Basic) Options(ctx Context) {
	ctx.Error(http.StatusMethodNotAllowed, ErrNotImplement)
}

// Trace Basic.Trace
func (b *Basic) Trace(ctx Context) {
	ctx.Error(http.StatusMethodNotAllowed, ErrNotImplement)
}

// Handle Basic.Handle
func (b *Basic) Handle(ctx Context) {
	ctx.Error(http.StatusMethodNotAllowed, ErrNotImplement)
}
