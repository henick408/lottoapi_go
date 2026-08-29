package main

import (
	"context"
	"fmt"
	"lottoapi/client"
	resultHandler "lottoapi/handler/results"
	resultService "lottoapi/service/results"

	"github.com/go-resty/resty/v2"
	"github.com/labstack/echo/v5"
)

var restyClient = resty.New()

func main() {

	e := echo.New()

	client := client.NewLottoClient(restyClient)
	resultService := resultService.NewResultService(client)
	resultHandler := resultHandler.NewResultHandler(resultService)
	e.GET("/results", resultHandler.GetAllHandler)
	//http.HandleFunc("GET /results", resultHandler.GetAllHandler)
	e.GET("/results/:gameType", resultHandler.GetByGameHandler)
	//http.HandleFunc("GET /results/{gameType}", resultHandler.GetByGameHandler)

	port := 8081
	fmt.Printf("Listening on port %d\n", port)
	sc := echo.StartConfig{Address: fmt.Sprintf(":%d", port)}
	if err := sc.Start(context.Background(), e); err != nil {
		e.Logger.Error("Failed to start server", "error", err)
	}

	// err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	// if err != nil {
	// 	panic(err)
	// }

}
