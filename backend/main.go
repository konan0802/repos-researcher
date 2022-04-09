package main

import (
	"repos-researcher/handler"
	"repos-researcher/infra"
	"repos-researcher/usecase"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// httpリクエストの情報を取得
	body := request.Body
	pathParam := request.PathParameters["pathparam"]
	queryParam := request.QueryStringParameters["queryparam"]

	// 依存関係を注入
	gInf := infra.NewGitHubInfra()
	dInf := infra.NewDynamoDBInfra()
	uc := usecase.NewUseCase(gInf, dInf)
	hdr := handler.NewHandler(uc)

	// ルーティングの設定
	var res []byte
	switch pathParam {
	case "/searchaccount":
		res = hdr.SearchAccount(queryParam)
	case "/searchrepository":
		res = hdr.SearchRepository(queryParam)
	case "/fetchaccount":
		res = hdr.FetchAccount(queryParam)
	case "/fetchrepository":
		res = hdr.FetchRepository(queryParam)
	case "/saveaccount":
		res = hdr.SaveAccount(body)
	case "/saverepository":
		res = hdr.SaveRepository(body)
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
