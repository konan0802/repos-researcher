package handler

import (
	"fmt"
	"repos-researcher/domain/model"
	"repos-researcher/usecase"

	"github.com/aws/aws-lambda-go/events"
)

type Handler interface {
	SearchAccount(events.APIGatewayProxyRequest) []byte
	SearchRepository(events.APIGatewayProxyRequest) []byte
	FetchAccount(events.APIGatewayProxyRequest) []byte
	FetchRepository(events.APIGatewayProxyRequest) []byte
	SaveAccount(events.APIGatewayProxyRequest) []byte
	SaveRepository(events.APIGatewayProxyRequest) []byte
}

type handler struct {
	usc usecase.Usecase
}

func NewHandler(usc usecase.Usecase) Handler {
	return &handler{
		usc: usc,
	}
}

func (hdr handler) SearchAccount(request events.APIGatewayProxyRequest) []byte {
	var query model.SearchAccountQueryForm
	test2 := request.QueryStringParameters["test"]
	fmt.Printf("test2: %v\n", test2)

	b := []byte("SearchAccountTest " + test2 + "??")
	return b
}

func (hdr handler) SearchRepository(request events.APIGatewayProxyRequest) []byte {
	b := []byte("SearchRepositoryTest")
	return b
}
func (hdr handler) FetchAccount(request events.APIGatewayProxyRequest) []byte {
	b := []byte("FetchAccountTest")
	return b
}
func (hdr handler) FetchRepository(request events.APIGatewayProxyRequest) []byte {
	b := []byte("FetchRepositoryTest")
	return b
}
func (hdr handler) SaveAccount(request events.APIGatewayProxyRequest) []byte {
	b := []byte("SaveAccountTest")
	return b
}
func (hdr handler) SaveRepository(request events.APIGatewayProxyRequest) []byte {
	b := []byte("SaveRepositoryTest")
	return b
}
