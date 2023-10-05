package server

import (
	"fmt"
	"log"
	"net/http"
)

type Config struct{}

func InitServer(webPort string) *http.Server {
	app := Config{}

	log.Printf("Starting broker service on port %s", webPort)

	// Define http server
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}

	return srv
}
