package app

import (
	"context"
	"net/http"

	"github.com/bwmarrin/discordgo"
	"github.com/google/go-github/v52/github"
	"github.com/merlinfuchs/dissue/db"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"golang.org/x/oauth2"
)

type App struct {
	db  *db.KVDatabase
	dis *discordgo.Session
	gh  *github.Client
}

func New(db *db.KVDatabase) (*App, error) {
	dis, err := discordgo.New("Bot " + viper.GetString("discord.token"))
	if err != nil {
		return nil, err
	}

	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: viper.GetString("github.token")},
	)
	tc := oauth2.NewClient(ctx, ts)

	gh := github.NewClient(tc)

	return &App{
		db:  db,
		dis: dis,
		gh:  gh,
	}, nil
}

func (app *App) Start() error {
	http.HandleFunc("/webhook", app.handleWebhook)

	app.dis.AddHandler(app.handleThreadCreated)

	err := app.dis.Open()
	if err != nil {
		return err
	}

	address := viper.GetString("api.host") + ":" + viper.GetString("api.port")
	log.Info().Msgf("listening on %s", address)
	return http.ListenAndServe(address, nil)
}
