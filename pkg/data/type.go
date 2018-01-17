package data

import "github.com/segmentio/ksuid"

type Type struct {
	ID   string                 `json:"id" dynamodbav:"Id"`
	Data []string `json:"data" dynamodbav:"Data"`
}

func New(data []string) *Type {
	return &Type{
		ID:   ksuid.New().String(),
		Data: data,
	}
}
