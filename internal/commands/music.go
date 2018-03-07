package commands

import (
	"fmt"
	"github.com/codeclimate/hestia/internal/music"
	"github.com/codeclimate/hestia/internal/notifiers"
	"github.com/codeclimate/hestia/internal/types"
	"strings"
)

type Music struct {
	User     string
	Input    types.Input
	Notifier notifiers.Notifier
}

func (c Music) Run() {
	var message string

	switch c.Input.Args {
	case "playlists":
		message = c.listPlaylists()
	case "state":
		message = c.getState()
	default:
		message = c.invokeCommand()
	}

	c.Notifier.Log(message)
}

func (c Music) listPlaylists() string {
	playlists := music.ListPlaylists()

	return fmt.Sprintf("<@%s>:\n%s", c.User, strings.Join(playlists, "\n"))
}

func (c Music) getState() string {
	state := music.GetState()

	return fmt.Sprintf("<@%s>: [%s / %d] %s by %s on %s", c.User, state.Status, state.Volume, state.Title, state.Artist, state.Album)
}

func (c Music) invokeCommand() string {
	parts := strings.Split(c.Input.Args, " ")
	resp := music.InvokeCommand(parts[0], strings.Join(parts[1:], " "))

	if resp.Error != "" {
		return fmt.Sprintf("<@%s>: %s", c.User, resp.Error)
	} else {
		return c.getState()
	}
}
