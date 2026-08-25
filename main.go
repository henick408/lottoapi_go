package main

import (
	"fmt"
	"lottoapi/client"
	resultHandler "lottoapi/handler/results"
	resultService "lottoapi/service/results"
	"net/http"
	"time"
)

var httpClient = &http.Client{
	Timeout: 10 * time.Second,
}

func main() {

	client := client.NewLottoClient(httpClient)
	resultService := resultService.NewResultService(client)
	resultHandler := resultHandler.NewResultHandler(resultService)
	http.HandleFunc("GET /results", resultHandler.GetAllHandler)
	http.HandleFunc("GET /results/{gameType}", resultHandler.GetByGameHandler)

	port := 8080
	fmt.Printf("Listening on port %d\n", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		panic(err)
	}

}
