package dynamodb

import (
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/jkaflik/aws-lambda-go-api-example/pkg/data"
)

type DataRepository struct {
	TableName string
	db        *dynamodb.DynamoDB
}

func NewDataRepository(tableName string, db *dynamodb.DynamoDB) *DataRepository {
	return &DataRepository{
		TableName: tableName,
		db:        db,
	}
}

func (repository *DataRepository) GetAll() ([]*data.Type, error) {
	response, err := repository.db.Scan(&dynamodb.ScanInput{
		TableName: &repository.TableName,
	})

	if err != nil {
		return nil, err
	}

	var list []*data.Type

	err = dynamodbattribute.UnmarshalListOfMaps(response.Items, &list)

	if err != nil {
		return nil, err
	}

	return list, nil
}

func (repository *DataRepository) Save(data *data.Type) error {
	item, err := dynamodbattribute.MarshalMap(data)

	if err != nil {
		return err
	}

	_, err = repository.db.PutItem(&dynamodb.PutItemInput{
		TableName: &repository.TableName,
		Item:      item,
	})

	return err
}
