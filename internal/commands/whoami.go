package commands

import (
	"fmt"
	"github.com/nlopes/slack"
	"log"
)

func WhoAmI(userId string, channel string, client *slack.Client) {
	user, err := client.GetUserInfo(userId)

	if err != nil {
		log.Fatal(err)
	}

	message := fmt.Sprintf("<@%s>:\n id: %s\n name: %s\n email: %s", user.ID, user.ID, user.Profile.RealName, user.Profile.Email)

	postParams := slack.PostMessageParameters{}
	_, _, err = client.PostMessage(channel, message, postParams)
}
