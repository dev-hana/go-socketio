package main

import (
	"log"

	"github.com/dev-hana/go-socketio/routers"
)

func main() {
	r := routers.RunAPIServer()
	log.Println("Serving at localhost:8080...")
	r.Run(":8080")
}
