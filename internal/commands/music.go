package commands

import (
	"fmt"
	"github.com/codeclimate/hestia/internal/config"
	"github.com/codeclimate/hestia/internal/music"
	"github.com/codeclimate/hestia/internal/notifiers"
	"github.com/codeclimate/hestia/internal/types"
	"strings"
	"time"
)

type Music struct {
	User     string
	Input    types.Input
	Notifier notifiers.Notifier
}

func (c Music) Run() {
	var message string

	switch c.Input.Args {
	case "info":
		message = c.info()
	case "playlists":
		message = c.listPlaylists()
	case "state":
		message = c.getState()
	default:
		message = c.invokeCommand()
	}

	c.Notifier.Log(message)
}

func (c Music) info() string {
	var lines []string
	app_store_url := "https://itunes.apple.com/us/app/volumio/id1268256519"
	play_store_url := "https://play.google.com/store/apps/details?id=volumio.browser.Volumio"

	lines = append(lines, fmt.Sprintf("web ui: %s", config.Fetch("music_domain")))
	lines = append(lines, fmt.Sprintf("  username: %s", config.Fetch("music_username")))
	lines = append(lines, fmt.Sprintf("apps: <%s|ios>, <%s|android>", app_store_url, play_store_url))

	return fmt.Sprintf("<@%s>:\n%s", c.User, strings.Join(lines, "\n"))
}

func (c Music) listPlaylists() string {
	playlists := music.ListPlaylists()

	return fmt.Sprintf("<@%s>:\n%s", c.User, strings.Join(playlists, "\n"))
}

func (c Music) getState() string {
	state := music.GetState()

	return fmt.Sprintf("[%s / %d] %s by %s on %s", state.Status, state.Volume, state.Title, state.Artist, state.Album)
}

func (c Music) invokeCommand() string {
	parts := strings.Split(c.Input.Args, " ")
	resp := music.InvokeCommand(parts[0], strings.Join(parts[1:], " "))

	if resp.Error != "" {
		return resp.Error
	} else {
		// wait a second for system to react
		time.Sleep(1 * time.Second)

		return c.getState()
	}
}

func (c Music) HelpText() string {
	return "music <subcommand> [args]"
}

func (c Music) HelpDescription() string {
	return "Controls the office music server"
}

func (c Music) HelpExamples() []string {
	return []string{
		"music state",
		"music info",
		"music play",
		"music stop",
		"music pause",
		"music next",
		"music prev",
		"music volume 80",
		"music volume +|-|mute|unmute",
		"music playplaylist <name>",
		"music clearQueue",
		"music playlists",
	}
}
