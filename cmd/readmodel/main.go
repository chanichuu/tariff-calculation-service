package main

import (
	"fmt"
	"tariff-calculation-service/internal/readmodel"
	"tariff-calculation-service/internal/router"
	"tariff-calculation-service/pkg"

	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	router := router.NewRouter()
	readmodel.RouteReadmodelCalls(router)

	lambda.Start(pkg.AdaptGinRouter(router))
	fmt.Println("Started lambda handler.")
}
