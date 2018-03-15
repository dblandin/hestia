package commands

import (
	"fmt"
	"github.com/codeclimate/hestia/internal/config"
	"github.com/codeclimate/hestia/internal/notifiers"
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
	client := slack.New(config.Fetch("slack_bot_token"))

	user, err := client.GetUserInfo(c.User)

	if err != nil {
		log.Fatal(err)
	}

	message := fmt.Sprintf("\nid: %s\n name: %s\n email: %s", user.ID, user.ID, user.Profile.RealName, user.Profile.Email)

	c.Notifier.Log(message)
}

func (c WhoAmI) HelpText() string {
	return "whoami"
}

func (c WhoAmI) HelpDescription() string {
	return "Returns information about your slack user"
}

func (c WhoAmI) HelpExamples() []string {
	return []string{"whoami"}
}
