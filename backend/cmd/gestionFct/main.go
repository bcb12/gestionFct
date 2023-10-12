package main

import (
	"log"
	"net/http"
	"os"

	"github.com/bcb12/gestionFct/api"
	"github.com/bcb12/gestionFct/internal/config"
)

const envFile = "config/.env.development"

func main() {
	config.LoadEnv(envFile)

	port, exists := os.LookupEnv("APP_PORT_DEV")
	if !exists {
		log.Fatal("The environment variable APP_PORT_DEV is not defined")
	}

	dsn, err := setUpDSN()
	if err != nil {
		log.Fatal("The data source name could not be determined")
	}

	conn := config.ConnectToDB(dsn)

	ctx := &config.AppContext{DB: conn}

	r := api.NewRouter(ctx)

	log.Printf("Starting server in port: %s...", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}

func setUpDSN() (string, error) {
	var dsn string
	parameters := [5][2]string{{"DB_USER", "user"}, {"DB_PASSWORD", "password"}, {"DB_DATABASE", "dbname"}, {"DB_ADDRESS", "host"}, {"DB_PORT", "port"}}

	for _, param := range parameters {
		value, exists := os.LookupEnv(param[0])
		if !exists {
			log.Fatalf("The environment variable %s is not defined", param)
		}
		dsn += param[1] + "=" + value + " "
	}

	log.Printf("DB connection params: %s", dsn)

	return dsn, nil
}
