package main

import (
	"os"

	"github.com/HIMASAKTA-DEV/himasakta-backend/cmd"
	"github.com/HIMASAKTA-DEV/himasakta-backend/core/config"
	"github.com/joho/godotenv"
)

func initEnv() {
	// Hanya load file .env kalau ada (dev)
	if _, err := os.Stat(".env"); err == nil {
		if err := godotenv.Load(); err != nil {
			panic("Failed to loading env file")
		}
	}
}

func main() {
	initEnv()

	if err := cmd.Commands(); err != nil {
		panic("Failed Get Commands: " + err.Error())
	}

	RestApi := config.NewRest()
	RestApi.Start()
}

