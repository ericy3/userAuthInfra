package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
	"lambda-func/app"
)

type MyEvent struct {
	Username string `json:"username"`
}

// Take in a payload and
func HandleRequest(event MyEvent) (string, error) {
	if event.Username == "" {
		return "", fmt.Errorf("username cannot be empty")
	}

	return fmt.Sprintf("Successfully called by - %s", event.Username), nil
}

func main() {
	app := app.NewApp()
	lambda.Start(HandleRequest)
}
