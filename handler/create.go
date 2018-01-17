package main

import (
	"context"

	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws/session"
	awsdynamodb "github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/jkaflik/aws-lambda-go-api-example/pkg/data"
	"github.com/jkaflik/aws-lambda-go-api-example/pkg/dynamodb"
	"github.com/jkaflik/aws-lambda-go-api-example/pkg/http"
	"os"
)

func main() {
	db := awsdynamodb.New(session.New())
	repo := dynamodb.NewDataRepository(os.Getenv("DYNAMODB_DATA_TABLE"), db)

	lambda.Start(func(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
		var body http.DataInput
		err := json.Unmarshal([]byte(request.Body), &body)

		if err != nil {
			return http.NewJSONErrorResponse(400, fmt.Errorf("Invalid data: %s", err)), nil
		}

		dat := data.New(body.Data)
		err = repo.Save(dat)

		if err != nil {
			return http.NewJSONErrorResponse(500, fmt.Errorf("Persistence error: %s", err)), err
		}

		return http.NewJSONResponse(200, dat), nil
	})
}
