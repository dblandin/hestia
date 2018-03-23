package commands

import (
	"github.com/codeclimate/hestia/internal/notifiers"
	"github.com/codeclimate/hestia/internal/types"
)

type Command interface {
	Run()
	HelpText() string
	HelpDescription() string
	HelpExamples() []string
}

func Build(user string, input types.Input, notifier notifiers.Notifier) Command {
	var command Command

	if input.Args == "help" {
		input := types.Input{Args: input.Command}
		command = Help{user, input, notifier}
	} else {
		switch input.Command {
		case "boom":
			command = Boom{user, input, notifier}
		case "danceparty":
			command = DanceParty{user, input, notifier}
		case "echo":
			command = Echo{user, input, notifier}
		case "help":
			command = Help{user, input, notifier}
		case "nowplaying":
			command = NowPlaying{user, input, notifier}
		case "music":
			command = Music{user, input, notifier}
		case "weather":
			command = Weather{user, input, notifier}
		case "whoami":
			command = WhoAmI{user, input, notifier}
		default:
			command = Fallback{user, input, notifier}
		}
	}

	return command
}

func allCommands() []Command {
	return []Command{
		new(Boom),
		new(DanceParty),
		new(Echo),
		new(Help),
		new(Music),
		new(NowPlaying),
		new(Weather),
		new(WhoAmI),
	}
}
