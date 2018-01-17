package main

import (
	"context"

	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws/session"
	awsdynamodb "github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/jkaflik/aws-lambda-go-api-example/pkg/dynamodb"
	"github.com/jkaflik/aws-lambda-go-api-example/pkg/http"
	"os"
)

func main() {
	db := awsdynamodb.New(session.New())
	repo := dynamodb.NewDataRepository(os.Getenv("DYNAMODB_DATA_TABLE"), db)

	lambda.Start(func(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
		dataList, err := repo.GetAll()

		if err != nil {
			return http.NewJSONErrorResponse(500, fmt.Errorf("Read error: %s", err)), err
		}

		return http.NewJSONResponse(200, dataList), nil
	})
}
