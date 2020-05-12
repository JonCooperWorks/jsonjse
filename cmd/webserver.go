package main

import (
	"log"
	"net/http"

	"github.com/joncooperworks/jsonjse"
)

func main() {
	jse := &jsonjse.JSE{
		Client: &http.Client{},
	}
	config := &jsonjse.ServerConfig{
		JSE: jse,
	}
	router := jsonjse.Router(config)
	err := router.Run()
	if err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
