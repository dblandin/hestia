package commands

import (
	"fmt"
	"github.com/nlopes/slack"
	"log"
)

func Default(userId string, channel string, command string, client *slack.Client) {
	message := fmt.Sprintf("<@%s> command `%s` not found", userId, command)

	postParams := slack.PostMessageParameters{}
	_, _, err := client.PostMessage(channel, message, postParams)

	if err != nil {
		log.Fatal(err)
	}
}
