package main

import (
	"repos-researcher/controller"
	"repos-researcher/infra"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	// 依存関係を注入
	gInf := infra.NewGitHubInfra()
	dInf := infra.NewDynamoDBInfra()
	hdr := controller.NewController(gInf, dInf)

	// ルーティングの設定
	var res []byte
	switch request.Resource {
	case "/searchaccount":
		res = hdr.SearchAccount(request)
	case "/searchrepository":
		res = hdr.SearchRepository(request)
	case "/fetchaccount":
		res = hdr.FetchAccount(request)
	case "/fetchrepository":
		res = hdr.FetchRepository(request)
	case "/saveaccount":
		res = hdr.SaveAccount(request)
	case "/saverepository":
		res = hdr.SaveRepository(request)
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
