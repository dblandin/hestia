package commands

import (
	"fmt"
	"github.com/nlopes/slack"
	"log"
)

func Echo(userId string, channel string, args string, client *slack.Client) {
	message := fmt.Sprintf("<@%s>: %s", userId, args)

	postParams := slack.PostMessageParameters{}
	_, _, err := client.PostMessage(channel, message, postParams)

	if err != nil {
		log.Fatal(err)
	}
}
