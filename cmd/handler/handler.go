package main

import (
	"context"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/codeclimate/hestia/internal/commands"
	"github.com/nlopes/slack"
	"log"
	"regexp"
)

type Response struct {
	Message string `json:"message"`
	Ok      bool   `json:"ok"`
}

type Event struct {
	Type           string `json:"type"`
	User           string `json:"user"`
	Text           string `json:"text"`
	Timestamp      string `json:"ts"`
	Channel        string `json:"channel"`
	EventTimestamp string `json:"event_ts"`
}

type EventCallback struct {
	Token       string   `json:"type"`
	Challenge   string   `json:"challenge"`
	TeamId      string   `json:"team_id"`
	ApiAppId    string   `json:"api_app_id"`
	Event       Event    `json:"event"`
	Type        string   `json:"type"`
	EventId     string   `json:"event_id"`
	EventTime   int      `json:"event_time"`
	AuthedUsers []string `json:"authed_users"`
}

// {
//     "token": "FeCDfP96MxGb3JA2TTmXVhmc",
//     "team_id": "T0G0RAKPG",
//     "api_app_id": "A9FDQB5V5",
//     "event": {
//         "type": "app_mention",
//         "user": "U0G0RMBC0",
//         "text": "<@U9EC5EG3U> sup?",
//         "ts": "1519519253.000080",
//         "channel": "C0G0KFXS8",
//         "event_ts": "1519519253000080"
//     },
//     "type": "event_callback",
//     "event_id": "Ev9E9CLT99",
//     "event_time": 1519519253000080,
//     "authed_users": [
//         "U9EC5EG3U"
//     ]
// }

func handleRequest(ctx context.Context, eventCallback EventCallback) (Response, error) {
	event := eventCallback.Event

	re := regexp.MustCompile(`<@\w+>\s+(?P<command>\w+)\s?(?P<args>.*)?`)
	match := re.FindStringSubmatch(event.Text)

	capturesMap := make(map[string]string)
	for i, name := range re.SubexpNames() {
		if i > 0 && i <= len(match) {
			capturesMap[name] = match[i]
		}
	}

	command := capturesMap["command"]
	args := capturesMap["args"]

	log.Printf("command = %s.\n", command)
	log.Printf("args = %s.\n", args)

	client := slack.New("abc123")

	switch command {
	case "whoami":
		commands.WhoAmI(event.User, event.Channel, client)
	case "echo":
		commands.Echo(event.User, event.Channel, args, client)
	default:
		commands.Default(event.User, event.Channel, command, client)
	}

	return Response{Message: "processed message", Ok: true}, nil
}

func main() {
	lambda.Start(handleRequest)
}
