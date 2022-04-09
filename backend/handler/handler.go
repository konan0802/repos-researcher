package handler

import (
	"repos-researcher/usecase"
)

type Handler interface {
	SearchAccount(queryParam string) []byte
	SearchRepository(queryParam string) []byte
	FetchAccount(queryParam string) []byte
	FetchRepository(queryParam string) []byte
	SaveAccount(body string) []byte
	SaveRepository(body string) []byte
}

type handler struct {
	uc usecase.Usecase
}

func NewHandler(uc usecase.Usecase) Handler {
	return &handler{
		uc: uc,
	}
}

func (hdr handler) SearchAccount(queryParam string) []byte {
	b := []byte(queryParam)
	return b
}

func (hdr handler) SearchRepository(queryParam string) []byte {
	b := []byte(queryParam)
	return b
}
func (hdr handler) FetchAccount(queryParam string) []byte {
	b := []byte(queryParam)
	return b
}
func (hdr handler) FetchRepository(queryParam string) []byte {
	b := []byte(queryParam)
	return b
}
func (hdr handler) SaveAccount(body string) []byte {
	b := []byte(body)
	return b
}
func (hdr handler) SaveRepository(body string) []byte {
	b := []byte(body)
	return b
}
