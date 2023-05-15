package main

import (
	"github.com/rs/zerolog/log"

	"github.com/merlinfuchs/dissue/app"
	"github.com/merlinfuchs/dissue/config"
	"github.com/merlinfuchs/dissue/db"
)

func main() {
	config.InitConfig()

	db, err := db.New()
	if err != nil {
		log.Fatal().Err(err).Msg("could not open database")
	}

	app, err := app.New(db)
	if err != nil {
		log.Fatal().Err(err).Msg("could not create app")
	}

	err = app.Start()
	if err != nil {
		log.Fatal().Err(err).Msg("could not start app")
	}
}
