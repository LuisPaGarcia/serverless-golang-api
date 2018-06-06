package main

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// Something Hey
type Something struct {
	Identif int
	Name    string
	Age     int
}

// Handler export
func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	ret := Something{
		Identif: 1,
		Name:    "Jane Doe",
		Age:     26}

	bodyj, _ := json.Marshal(ret)
	output := string(bodyj)

	return events.APIGatewayProxyResponse{Body: output, StatusCode: 200}, nil
}

func main() {
	lambda.Start(Handler)
}
