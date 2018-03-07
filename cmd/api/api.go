package main

import (
	"context"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/bugsnag/bugsnag-go"
	"github.com/codeclimate/hestia/internal/config"
	"log"
	"os"
)

import awsLambda "github.com/aws/aws-sdk-go/service/lambda"

func init() {
	api_key := config.Fetch("bugsnag_api_key")

	bugsnag.Configure(bugsnag.Configuration{
		APIKey:          api_key,
		ReleaseStage:    os.Getenv("BUGSNAG_RELEASE_STAGE"),
		ProjectPackages: []string{"github.com/codeclimate/hestia"},
		Synchronous:     true,
	})
}

func main() {
	lambda.Start(handleRequest)
}

func handleRequest(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	defer bugsnag.AutoNotify()

	log.Printf("Processing request data for request %s.\n", request.RequestContext.RequestID)

	service := awsLambda.New(session.New())

	input := &awsLambda.InvokeInput{
		FunctionName:   aws.String(os.Getenv("COMMAND_LAMBDA_FUNCTION_NAME")),
		InvocationType: aws.String("Event"),
		Payload:        []byte(request.Body),
		Qualifier:      aws.String(os.Getenv("COMMAND_LAMBDA_VERSION")),
	}

	_, err := service.Invoke(input)

	if err != nil {
		log.Fatal(err)
	}

	return events.APIGatewayProxyResponse{Body: request.Body, StatusCode: 200}, nil
}
