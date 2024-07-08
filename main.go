package main

import (
	"log"

	"github.com/Moxeed5/dockerPractice/config"
	"github.com/Moxeed5/dockerPractice/internal/db"
)

func main() {
	cfg := config.LoadConfig()
	db.Init(cfg)
	log.Println("app started")
}
