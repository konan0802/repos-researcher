package controller

import (
	"fmt"
	"repos-researcher/domain/repository"

	"github.com/aws/aws-lambda-go/events"
)

type Controller interface {
	SearchAccount(events.APIGatewayProxyRequest) []byte
	SearchRepository(events.APIGatewayProxyRequest) []byte
	FetchAccount(events.APIGatewayProxyRequest) []byte
	FetchRepository(events.APIGatewayProxyRequest) []byte
	SaveAccount(events.APIGatewayProxyRequest) []byte
	SaveRepository(events.APIGatewayProxyRequest) []byte
}

type controller struct {
	grp repository.GitHubRepository
	drp repository.DynamoDBRepository
}

func NewController(grp repository.GitHubRepository, drp repository.DynamoDBRepository) Controller {
	return &controller{
		grp: grp,
		drp: drp,
	}
}

func (hdr controller) SearchAccount(request events.APIGatewayProxyRequest) []byte {
	test2 := request.QueryStringParameters["test"]
	fmt.Printf("test2: %v\n", test2)

	b := []byte("SearchAccountTest " + test2 + "??")
	return b
}

func (hdr controller) SearchRepository(request events.APIGatewayProxyRequest) []byte {
	b := []byte("SearchRepositoryTest")
	return b
}
func (hdr controller) FetchAccount(request events.APIGatewayProxyRequest) []byte {
	b := []byte("FetchAccountTest")
	return b
}
func (hdr controller) FetchRepository(request events.APIGatewayProxyRequest) []byte {
	b := []byte("FetchRepositoryTest")
	return b
}
func (hdr controller) SaveAccount(request events.APIGatewayProxyRequest) []byte {
	b := []byte("SaveAccountTest")
	return b
}
func (hdr controller) SaveRepository(request events.APIGatewayProxyRequest) []byte {
	b := []byte("SaveRepositoryTest")
	return b
}
