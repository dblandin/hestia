package commands

import (
	"fmt"
	"github.com/codeclimate/hestia/internal/types"
	"github.com/nlopes/slack"
	"github.com/shkh/lastfm-go/lastfm"
	"log"
	"strings"
)

type NowPlaying struct {
	Event  types.Event
	Input  types.Input
	Client *slack.Client
}

func (command NowPlaying) Run() {
	api := lastfm.New("abc123", "abc123")

	usernames := []string{"infinitedevon"}
	var output []string

	for _, username := range usernames {
		result, _ := api.User.GetRecentTracks(lastfm.P{
			"user":  username,
			"limit": 1,
		})

		recentTrack := result.Tracks[0]

		if len(recentTrack.NowPlaying) > 0 {
			output = append(output, fmt.Sprintf("%s by %s (%s)", recentTrack.Name, recentTrack.Artist.Name, username))
		}
	}

	var message string

	if len(output) > 0 {
		message = fmt.Sprintf("<@%s>: now playing\n%s", command.Event.User, strings.Join(output, "\n"))
	} else {
		message = fmt.Sprintf("<@%s>: all quiet", command.Event.User)
	}

	postParams := slack.PostMessageParameters{}
	_, _, err := command.Client.PostMessage(command.Event.Channel, message, postParams)

	if err != nil {
		log.Fatal(err)
	}
}
