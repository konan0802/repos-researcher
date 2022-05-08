package main

import (
	"repos-researcher/handler"
	"repos-researcher/infra"
	"repos-researcher/usecase"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	// 依存関係を注入
	gInf := infra.NewGitHubInfra()
	dInf := infra.NewDynamoDBInfra()
	usc := usecase.NewUsecase(gInf, dInf)
	hdr := handler.NewHandler(usc)

	// ルーティングの設定
	var res []byte
	switch request.Resource {
	case "/search/account":
		res = hdr.SearchAccount(request)
	case "/search/repository":
		res = hdr.SearchRepository(request)
	case "/save/account":
		res = hdr.SaveAccount(request)
	case "/save/repository":
		res = hdr.SaveRepository(request)
	case "/fetch/savedaccount":
		res = hdr.FetchSavedAccount(request)
	case "/fetch/savedrepository":
		res = hdr.FetchSavedRepository(request)
	case "/delete/savedaccount":
		res = hdr.DeleteSavedAccount(request)
	case "/delete/savedrepository":
		res = hdr.DeleteSavedRepository(request)
	default:
		return events.APIGatewayProxyResponse{
			Body:       `Error: 404 not found`,
			StatusCode: 404,
		}, nil
	}

	return events.APIGatewayProxyResponse{
		Body:       string(res),
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(Handler)
}
