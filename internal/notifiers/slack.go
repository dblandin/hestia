package notifiers

import (
	"github.com/codeclimate/hestia/internal/config"
	"github.com/nlopes/slack"
	"log"
)

type Slack struct {
	Channel string
}

func (n Slack) Log(message string) {
	client := slack.New(config.Fetch("slack_bot_token"))

	postParams := slack.PostMessageParameters{}
	_, _, err := client.PostMessage(n.Channel, message, postParams)

	if err != nil {
		log.Fatal(err)
	}
}
