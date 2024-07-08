package db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/Moxeed5/dockerPractice/config"
	_ "github.com/lib/pq" // Import the PostgreSQL driver
)

var DB *sql.DB

// Config is a struct type from the config package. Using dot notation here is needed because we're using Config in another package.
func Init(cfg config.Config) {
	var err error

	//using Sprintf here bc it returns a string whereas printf just prints to console.
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost, cfg.DBPort, cfg.DBUsername, cfg.DBPassword, cfg.DBName)

	DB, err = sql.Open("postgres", dsn)

	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	log.Println("Successful connection to DB")

}
