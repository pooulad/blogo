package main

import (
	"fmt"
	"log"

	"github.com/pooulad/blogo/api"
	"github.com/pooulad/blogo/internal/app"
	"github.com/pooulad/blogo/internal/config"
	"github.com/pooulad/blogo/internal/database"
	"github.com/pooulad/blogo/utilities"
)

func runServer() {
	// log layer: should add later
	log.SetPrefix(fmt.Sprintf("%s --> ERORR: ", utilities.ColorizeMessage(utilities.ColorYellow, "blogo")))
	log.SetFlags(0)

	// config layer: setup config of project
	cfg, err := config.New()
	if err != nil {
		log.Fatal(err)
	}

	// store layer: connect DB and use
	store, err := database.Connect(cfg)
	if err != nil {
		log.Fatal(err)
	}

	// application layer: handle logic of program
	app := app.New(store, cfg)

	// http/api layer: handle http/api requests
	api := api.New(app)
	log.Fatal(api.Start())
}
