package http

import (
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
)

func NewJSONResponse(statusCode int, body interface{}) events.APIGatewayProxyResponse {
	bodyMarshaled, err := json.Marshal(body)

	if err != nil {
		panic(err)
	}

	return events.APIGatewayProxyResponse{
		StatusCode: statusCode,
		Body:       string(bodyMarshaled),
	}
}

func NewJSONErrorResponse(statusCode int, err error) events.APIGatewayProxyResponse {
	return NewJSONResponse(statusCode, ErrorResponse{ err.Error(), statusCode })
}

type ErrorResponse struct {
	Error string `json:"error"`
	StatusCode int `json:"status_code"`
}