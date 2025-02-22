package main

import (
    "context"
    "github.com/aws/aws-lambda-go/lambda"
    "github.com/aws/aws-lambda-go/events"
    "github.com/LinhNguyen2901/url-classification/back-end/internal/api"
    "log"
)

func router(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
    log.Println("Received request")
    handler, err := api.NewHandler()
    if err != nil {
        return events.APIGatewayProxyResponse{
            StatusCode: 500,
            Body:       err.Error(),
        }, err
    }

	switch request.Path {
        case "/classify":
            log.Println("classify")
            return handler.HandleClassify(ctx, request)
        default:
            return events.APIGatewayProxyResponse{
                StatusCode: 404,
                Body:       "Not Found",
            }, nil
	}
}

func main() {
    lambda.Start(router)
}