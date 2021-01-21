package db

import (
	"github.com/go-pg/pg"
	"log"
	"os"
)

func Connect()  {
	opts := &pg.Options {
		User: "",
		Password: "",
		Addr: "",
		Database: "",
	}

	var db *pg.DB = pg.Connect(opts)
	if db == nil {
		log.Printf("Database connection failed.\n")
		os.Exit(100)
	}
	log.Printf("Connect successful.\n")
	closeErr := db.Close()
	if closeErr != nil {
		log.Printf("Error while closing connection, %v\n", closeErr)
		os.Exit(100)
	}
	log.Printf("Connection closed successfully.\n")
	return
}
