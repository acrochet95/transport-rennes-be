package main

import (
	"os"

	skill "github.com/acrochet95/transport-rennes-be/internal/alexa-skill"
	"github.com/aws/aws-lambda-go/lambda"
	"gitlab.com/dasjott/alexa-sdk-go"
)

func main() {
	skill.Initialize()

	alexa.AppID = os.Getenv("ALEXA_APP_ID")
	alexa.Handlers = skill.Handlers
	alexa.LocaleStrings = skill.Locales
	lambda.Start(alexa.Handle)
}
