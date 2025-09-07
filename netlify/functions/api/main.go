package main

import (
	"context"

	"github.com/aws/aws-lambda-go/lambda"
	echoadapter "github.com/awslabs/aws-lambda-go-api-proxy/echo"

	"heloo-go/internal/app"
)

var echoLambda *echoadapter.EchoLambda

func init() {
	server, _, err := app.NewServer()
	if err != nil {
		panic(err)
	}
	echoLambda = echoadapter.New(server)
}

func main() {
	lambda.StartWithOptions(echoLambda.ProxyWithContext, lambda.WithContext(context.Background()))
}
