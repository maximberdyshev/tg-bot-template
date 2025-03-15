package main

import (
	"log"

	"tg-bot-template/config"
	"tg-bot-template/internal/app"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	app, err := app.New(cfg)
	if err != nil {
		log.Fatal(err)
	}

	app.Run()
}
