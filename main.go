package main

import (
	"basic_api/client"
	resultHandler "basic_api/handler/results"
	resultService "basic_api/service/results"
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
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}

}
