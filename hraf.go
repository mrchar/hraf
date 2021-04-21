package hraf

import (
	"net/http"

	def "github.com/mrchar/hraf/definition"
	"github.com/mrchar/hraf/implement/server"
)

func New(address string, router def.Router) def.Server {
	return server.New(address, router)
}

func Default() def.Server {
	mux := http.NewServeMux()
	return server.New("localhost:8080", mux)
}
