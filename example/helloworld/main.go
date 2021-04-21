package main

import (
	"log"
	"net/http"

	"github.com/mrchar/hraf"
)

func main() {
	server := hraf.Default()

	server.HandleFunc("/greet", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`{"message":"hello world!"}`))
	})

	if err := server.Start(); err != nil {
		log.Fatal(err)
	}
}
