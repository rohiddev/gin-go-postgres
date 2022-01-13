package config

import (
	"log"
	"os"
	"github.com/go-pg/pg/v9"

	controllers "go-gin-postgres/controllers"
)

// Connecting to db
func Connect() *pg.DB {
	opts := &pg.Options{
		User: "rohiddev",
		Password: "****",
		Addr: "127.0.0.1:5432",
		Database: "rohiddev",
	}

	var db *pg.DB = pg.Connect(opts)
	if db == nil {
		log.Printf("Failed to connect")
		os.Exit(100)
	}
	log.Printf("Connected to db")
	controllers.CreateTodoTable(db)
	controllers.InitiateDB(db)
	return db
}
