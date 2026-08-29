package main

import (
	"fmt"
	"lottoapi/client"
	resultHandler "lottoapi/handler/results"
	resultService "lottoapi/service/results"
	"net/http"

	"github.com/go-resty/resty/v2"
)

var restyClient = resty.New()

func main() {

	client := client.NewLottoClient(restyClient)
	resultService := resultService.NewResultService(client)
	resultHandler := resultHandler.NewResultHandler(resultService)
	http.HandleFunc("GET /results", resultHandler.GetAllHandler)
	http.HandleFunc("GET /results/{gameType}", resultHandler.GetByGameHandler)

	port := 8081
	fmt.Printf("Listening on port %d\n", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		panic(err)
	}

}
