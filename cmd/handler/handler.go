package main

import (
	"context"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/bugsnag/bugsnag-go"
	"github.com/codeclimate/hestia/internal/commands"
	"github.com/codeclimate/hestia/internal/config"
	"github.com/codeclimate/hestia/internal/notifiers"
	"github.com/codeclimate/hestia/internal/types"
	"github.com/codeclimate/hestia/internal/utils"
	"log"
	"os"
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
	api_key := config.Fetch("bugsnag_api_key")

	bugsnag.Configure(bugsnag.Configuration{
		APIKey:          api_key,
		ReleaseStage:    os.Getenv("BUGSNAG_RELEASE_STAGE"),
		ProjectPackages: []string{"github.com/codeclimate/hestia"},
	})

	event := eventCallback.Event

	re := regexp.MustCompile(`<@\w+>\s+(?P<command>\w+)\s?(?P<args>.*)?`)
	input := utils.ExtractInput(event.Text, re)

	log.Printf("command = %s.\n", input.Command)
	log.Printf("args = %s.\n", input.Args)

	notifier := notifiers.Slack{Channel: event.Channel}

	command := commands.Build(event.User, input, notifier)
	command.Run()

	return Response{Message: "Processed message", Ok: true}, nil
}
