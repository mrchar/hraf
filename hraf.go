package hraf

import (
	"io/ioutil"
	"net/http"

	def "github.com/mrchar/hraf/definition"
	"github.com/mrchar/hraf/implement/codec/json"
	"github.com/mrchar/hraf/implement/router"
	"github.com/mrchar/hraf/implement/server"
)

func New(address string, router def.Router) def.Server {
	return server.New(address, router)
}

func Default(address string) def.Server {
	return server.New(address, router.New())
}

func Pluck(r *http.Request, v interface{}) error {
	buffer, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	if err = json.Decode(buffer, v); err != nil {
		return err
	}

	return nil
}

func Respond(w http.ResponseWriter, v interface{}) error {
	w.WriteHeader(http.StatusOK)

	buffer, err := json.Encode(v)
	if err != nil {
		return err
	}
	_, err = w.Write(buffer)
	return err
}
