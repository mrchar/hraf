package main

import (
	"log"
	"net/http"

	"github.com/mrchar/hraf"
)

func main() {
	server := hraf.Default(":8080")

	server.HandleFunc("/greet", greet)

	if err := server.Start(); err != nil {
		log.Fatal(err)
	}
}

type greetParams struct {
	Name string `json:"name"`
}

type greetResp struct {
	Message string `json:"message"`
}

func greet(w http.ResponseWriter, r *http.Request) {
	var params greetParams

	if err := hraf.Pluck(r, &params); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := hraf.Respond(w, greetResp{"hello " + params.Name + "!"}); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
