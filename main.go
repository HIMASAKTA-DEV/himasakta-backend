package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/HIMASAKTA-DEV/himasakta-backend/cmd"
	"github.com/HIMASAKTA-DEV/himasakta-backend/core/config"
	"github.com/joho/godotenv"
)

func initEnv() {
	// Hanya load file .env kalau ada (dev)
	if _, err := os.Stat(".env"); err == nil {
		path, _ := filepath.Abs(".env")
		if err := godotenv.Load(path); err != nil {
			log.Printf("Warning: Failed to loading env file at %s: %v", path, err)
		} else {
			log.Printf("Success: Loaded env file from %s", path)
			purl := os.Getenv("POSTGRES_URL")
			if purl != "" {
				log.Printf("POSTGRES_URL length: %d", len(purl))
			} else {
				log.Printf("POSTGRES_URL is EMPTY!")
			}
		}
	}
}

func main() {
	initEnv()

	if err := cmd.Commands(); err != nil {
		panic("Failed Get Commands: " + err.Error())
	}

	RestApi, err := config.NewRest()
	if err != nil {
		log.Fatalf("Failed to initialize REST API: %v", err)
	}
	RestApi.Start()
}
