package commands

import (
	"fmt"
	"github.com/codeclimate/hestia/internal/types"
	"github.com/nlopes/slack"
	"log"
)

type Fallback struct {
	Event  types.Event
	Input  types.Input
	Client *slack.Client
}

func (c Fallback) Run() {
	message := fmt.Sprintf("<@%s> command `%s` not found", c.Event.User, c.Input.Command)

	postParams := slack.PostMessageParameters{}
	_, _, err := c.Client.PostMessage(c.Event.Channel, message, postParams)

	if err != nil {
		log.Fatal(err)
	}
}
