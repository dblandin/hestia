package commands

import (
	"fmt"
	"github.com/codeclimate/hestia/internal/types"
	"github.com/nlopes/slack"
	"log"
)

type WhoAmI struct {
	Event  types.Event
	Input  types.Input
	Client *slack.Client
}

func (command WhoAmI) Run() {
	user, err := command.Client.GetUserInfo(command.Event.User)

	if err != nil {
		log.Fatal(err)
	}

	message := fmt.Sprintf("<@%s>:\n id: %s\n name: %s\n email: %s", user.ID, user.ID, user.Profile.RealName, user.Profile.Email)

	postParams := slack.PostMessageParameters{}
	_, _, err = command.Client.PostMessage(command.Event.Channel, message, postParams)
}
