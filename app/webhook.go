package app

import (
	"fmt"
	"net/http"

	"github.com/google/go-github/v52/github"
	"github.com/spf13/viper"
)

func (app *App) handleWebhook(w http.ResponseWriter, r *http.Request) {
	payload, err := github.ValidatePayload(r, []byte(viper.GetString("github.webhook_secret")))
	if err != nil {
	}
	event, err := github.ParseWebHook(github.WebHookType(r), payload)
	if err != nil {
	}
	switch event := event.(type) {
	case *github.IssueEvent:
		fmt.Println(event)
	case *github.IssueCommentEvent:
		fmt.Println(event)
	}
}
