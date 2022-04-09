package infra

import "repos-researcher/domain/repository"

type dynamoDBInfra struct{}

func NewDynamoDBInfra() repository.DynamoDBRepository {
	return &dynamoDBInfra{}
}
