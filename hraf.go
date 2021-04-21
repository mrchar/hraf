package hraf

import (
	def "github.com/mrchar/hraf/definition"
	"github.com/mrchar/hraf/implement/router"
	"github.com/mrchar/hraf/implement/server"
)

func New(address string, router def.Router) def.Server {
	return server.New(address, router)
}

func Default(address string) def.Server {
	return server.New(address, router.New())
}
