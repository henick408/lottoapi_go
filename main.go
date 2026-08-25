package main

import (
	"basic_api/client"
	resultHandler "basic_api/handler/results"
	resultService "basic_api/service/results"
	"fmt"
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

	port := 8080
	fmt.Printf("Listening on port %d\n", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		panic(err)
	}

}
