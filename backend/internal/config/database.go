package config

import (
	"database/sql"
	"log"
	"time"
)

type AppContext struct {
	DB *sql.DB
}

var counts int

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)

	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func ConnectToDB(dsn string) *sql.DB {

	for {
		connection, err := openDB(dsn)
		if err != nil {
			log.Printf("Postgres not yet ready: %s\n", err)
			counts++
		} else {
			log.Println("Connected to Postgres!")
			return connection
		}

		if counts > 10 {
			log.Println(err)
			return nil
		}

		log.Println("Backing off for two seconds...")
		time.Sleep(2 * time.Second)
		continue
	}
}
