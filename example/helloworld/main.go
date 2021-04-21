package main

import (
	"log"
	"net/http"

	"github.com/mrchar/hraf"
	"github.com/mrchar/hraf/implement/codec/json"
)

func main() {
	server := hraf.Default(":8080")

	server.HandleFunc("/greet", greet)

	if err := server.Start(); err != nil {
		log.Fatal(err)
	}
}

type greetResp struct {
	Message string `json:"message"`
}

func greet(w http.ResponseWriter, r *http.Request) {
	buffer, err := json.Encode(greetResp{"hello world!"})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(buffer)
}
