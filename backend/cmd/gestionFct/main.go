package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/bcb12/gestionFct/internal/config"
	server "github.com/bcb12/gestionFct/internal/http_server"
)

const envFile = "config/.env.development"

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world!")
}

func main() {
	config.LoadEnv(envFile)

	port, exists := os.LookupEnv("APP_PORT_DEV")
	if !exists {
		log.Fatal("The environment variable APP_PORT_DEV is not defined")
	}

	srv := server.InitServer(port)

	// Start the server
	err := srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
	http.HandleFunc("/", handler)
	log.Printf("Starting server in port: %s...", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
