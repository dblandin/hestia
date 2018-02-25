package main

import (
	"context"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/codeclimate/hestia/internal/commands"
	"github.com/codeclimate/hestia/internal/types"
	"github.com/nlopes/slack"
	"log"
	"regexp"
)

type Response struct {
	Message string `json:"message"`
	Ok      bool   `json:"ok"`
}

func main() {
	lambda.Start(handleRequest)
}

func handleRequest(ctx context.Context, eventCallback types.EventCallback) (Response, error) {
	event := eventCallback.Event
	input := extractInput(event.Text)

	log.Printf("command = %s.\n", input.Command)
	log.Printf("args = %s.\n", input.Args)

	client := slack.New("xoxb-320413492130-GgNGrajOmQeyCdQZrbzflqfa")

	command := commands.Build(event, input, client)
	command.Run()

	return Response{Message: "Processed message", Ok: true}, nil
}

func extractInput(text string) types.Input {
	re := regexp.MustCompile(`<@\w+>\s+(?P<command>\w+)\s?(?P<args>.*)?`)
	match := re.FindStringSubmatch(text)
	captures := extractCaptures(re, match)

	return types.Input{Command: captures["command"], Args: captures["args"]}
}

func extractCaptures(re *regexp.Regexp, match []string) map[string]string {
	captures := make(map[string]string)
	for i, name := range re.SubexpNames() {
		if i > 0 && i <= len(match) {
			captures[name] = match[i]
		}
	}

	return captures
}
