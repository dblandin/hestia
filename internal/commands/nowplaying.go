package commands

import (
	"fmt"
	"github.com/codeclimate/hestia/internal/notifiers"
	"github.com/codeclimate/hestia/internal/secrets"
	"github.com/codeclimate/hestia/internal/types"
	"github.com/shkh/lastfm-go/lastfm"
	"log"
	"strings"
)

type NowPlaying struct {
	User     string
	Input    types.Input
	Notifier notifiers.Notifier
}

func (c NowPlaying) Run() {
	api := lastfm.New(secrets.GetSecretValue("lastfm_api_key"), secrets.GetSecretValue("lastfm_api_secret"))

	usernames := []string{"infinitedevon"}
	var output []string

	for _, username := range usernames {
		result, err := api.User.GetRecentTracks(lastfm.P{
			"user":  username,
			"limit": 1,
		})

		if err != nil {
			log.Fatal(err)
		}

		recentTrack := result.Tracks[0]

		if len(recentTrack.NowPlaying) > 0 {
			output = append(output, fmt.Sprintf("%s by %s (%s)", recentTrack.Name, recentTrack.Artist.Name, username))
		}
	}

	var message string

	if len(output) > 0 {
		message = fmt.Sprintf("<@%s>: now playing\n%s", c.User, strings.Join(output, "\n"))
	} else {
		message = fmt.Sprintf("<@%s>: all quiet", c.User)
	}

	c.Notifier.Log(message)
}
