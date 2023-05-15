package app

import (
	"context"

	"github.com/bwmarrin/discordgo"
	"github.com/google/go-github/v52/github"
	"github.com/rs/zerolog/log"
)

func (app *App) handleThreadCreated(s *discordgo.Session, e *discordgo.ThreadCreate) {
	app.gh.Issues.Create(context.TODO(), "", "", &github.IssueRequest{
		Title: &e.Name,
		Body:  &e.Name,
	})
}

func (app *App) handleThreadDeleted(s *discordgo.Session, e *discordgo.ThreadDelete) {
	issueID, err := app.db.GetIssueIDForThreadID(e.ID)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get issue id for thread id")
		return
	}

	_, _, err = app.gh.Issues.Edit(context.TODO(), "", "", issueID, &github.IssueRequest{
		State: github.String("closed"),
	})
	if err != nil {
		log.Error().Err(err).Msg("Failed to close issue")
		return
	}
}

func (app *App) handleThreadUpdated(s *discordgo.Session, e *discordgo.ThreadUpdate) {}
