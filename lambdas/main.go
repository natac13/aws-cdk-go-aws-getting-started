package main

import (
	"encoding/json"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type App struct {
	id string
}

func newApp(id string) *App {

	return &App{id: id}
}

func (a *App) Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	responseBody := map[string]string{
		"message": "Hello, World!",
	}

	responseJSON, err := json.Marshal(responseBody)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Headers:    map[string]string{"Content-Type": "application/json"},
			Body:       `{"message": "Internal Server Error"}`,
		}, err
	}

	response := events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       string(responseJSON),
		Headers: map[string]string{
			"Content-Type":                     "text/plain",
			"Access-Control-Allow-Origin":      "*",
			"Access-Control-Allow-Methods":     "GET, POST, PUT, DELETE, OPTIONS",
			"Access-Control-Allow-Headers":     "Content-Type, Authorization, X-Requested-With",
			"Access-Control-Allow-Credentials": "true",
		},
	}

	return response, nil
}

func main() {

	id := "1234"

	app := newApp(id)

	lambda.Start(app.Handler)

}
