package handler

import (
	"fmt"
	"repos-researcher/usecase"

	"github.com/aws/aws-lambda-go/events"
)

type Handler interface {
	SearchAccount(events.APIGatewayProxyRequest) []byte
	SearchRepository(events.APIGatewayProxyRequest) []byte
	SaveAccount(events.APIGatewayProxyRequest) []byte
	SaveRepository(events.APIGatewayProxyRequest) []byte
	FetchSavedAccount(events.APIGatewayProxyRequest) []byte
	FetchSavedRepository(events.APIGatewayProxyRequest) []byte
	DeleteSavedAccount(events.APIGatewayProxyRequest) []byte
	DeleteSavedRepository(events.APIGatewayProxyRequest) []byte
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
	//svar query model.SearchAccountQueryForm
	test2 := request.QueryStringParameters["test"]
	fmt.Printf("test2: %v\n", test2)

	b := []byte("SearchAccountTest " + test2 + "??")
	return b
}

func (hdr handler) SearchRepository(request events.APIGatewayProxyRequest) []byte {
	b := []byte("SearchRepositoryTest")
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
func (hdr handler) FetchSavedAccount(request events.APIGatewayProxyRequest) []byte {
	b := []byte("FetchSavedAccountTest")
	return b
}
func (hdr handler) FetchSavedRepository(request events.APIGatewayProxyRequest) []byte {
	b := []byte("FetchSavedRepositoryTest")
	return b
}
func (hdr handler) DeleteSavedAccount(request events.APIGatewayProxyRequest) []byte {
	b := []byte("DeleteSavedAccountTest")
	return b
}
func (hdr handler) DeleteSavedRepository(request events.APIGatewayProxyRequest) []byte {
	b := []byte("DeleteSavedRepositoryTest")
	return b
}
