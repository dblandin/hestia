package main

import (
	"context"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"log"
	"os"
)

import awsLambda "github.com/aws/aws-sdk-go/service/lambda"

func handleRequest(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
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

func main() {
	lambda.Start(handleRequest)
}
