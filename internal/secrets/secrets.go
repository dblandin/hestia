package secrets

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"
	"log"
)

func GetSecretValue(key string) string {

	service := ssm.New(session.New())

	input := ssm.GetParameterInput{
		Name:           aws.String(fmt.Sprintf("hestia_production_handler.%s", key)),
		WithDecryption: aws.Bool(true),
	}

	output, err := service.GetParameter(&input)

	if err != nil {
		log.Fatal(err)
	}

	return *output.Parameter.Value
}
