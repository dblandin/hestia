package commands

import (
	"github.com/codeclimate/hestia/internal/types"
	"github.com/nlopes/slack"
)

type Command interface {
	Run()
}

func Build(event types.Event, input types.Input, client *slack.Client) Command {
	var command Command

	switch input.Command {
	case "whoami":
		command = WhoAmI{event, input, client}
	case "echo":
		command = Echo{event, input, client}
	case "nowplaying":
		command = NowPlaying{event, input, client}
	case "weather":
		command = Weather{event, input, client}
	default:
		command = Fallback{event, input, client}
	}

	return command
}
