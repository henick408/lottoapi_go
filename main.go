package main

import (
	"basic_api/client"
	"encoding/json"
	"net/http"
	"time"
)

var httpClient = &http.Client{
	Timeout: 10 * time.Second,
}

func getAllHandler(writer http.ResponseWriter, request *http.Request) {
	lottoClient := client.NewLottoClient(httpClient)
	results, err := lottoClient.GetLastResults()
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
	}

	writer.Header().Set("Content-Type", "application/json")

	json.NewEncoder(writer).Encode(results)

}

func main() {

	http.HandleFunc("GET /results", getAllHandler)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}

}
