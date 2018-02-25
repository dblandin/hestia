package commands

import (
	"fmt"
	"github.com/codeclimate/hestia/internal/notifiers"
	"github.com/codeclimate/hestia/internal/secrets"
	"github.com/codeclimate/hestia/internal/types"
	"github.com/nlopes/slack"
	"log"
)

type WhoAmI struct {
	User     string
	Input    types.Input
	Notifier notifiers.Notifier
}

func (c WhoAmI) Run() {
	client := slack.New(secrets.GetSecretValue("slack_bot_token"))

	user, err := client.GetUserInfo(c.User)

	if err != nil {
		log.Fatal(err)
	}

	message := fmt.Sprintf("<@%s>:\n id: %s\n name: %s\n email: %s", user.ID, user.ID, user.Profile.RealName, user.Profile.Email)

	c.Notifier.Log(message)
}
