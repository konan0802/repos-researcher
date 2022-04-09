package handler

import (
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
	uc usecase.Usecase
}

func NewHandler(uc usecase.Usecase) Handler {
	return &handler{
		uc: uc,
	}
}

func (hdr handler) SearchAccount(request events.APIGatewayProxyRequest) []byte {
	b := []byte("SearchAccountTest")
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
