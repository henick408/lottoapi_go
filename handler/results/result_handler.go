package results

import (
	"basic_api/model"
	"basic_api/response"
	"basic_api/service/results"
	"basic_api/util"
	"encoding/json"
	"net/http"
	"time"
)

type ResultHandler struct {
	service *results.ResultService
}

func NewResultHandler(service *results.ResultService) *ResultHandler {
	return &ResultHandler{
		service: service,
	}
}

func (handler *ResultHandler) GetAllHandler(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	drawDateStr := request.URL.Query().Get("drawDate")

	var (
		results []model.DrawResult
		err     error
	)

	if drawDateStr != "" {
		date, parseErr := time.Parse("2006-01-02", drawDateStr)
		if parseErr != nil {
			http.Error(writer, "Incorrect date", http.StatusBadRequest)
			return
		}
		drawDate := util.NewDate(
			date.Year(),
			int(date.Month()),
			date.Day(),
		)

		results, err = handler.service.GetResultsByDate(drawDate)
	} else {
		results, err = handler.service.GetLastResults()
	}

	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	var responses []response.ResultResponse

	for _, result := range results {
		responses = append(responses, result.ToResponse())
	}

	json.NewEncoder(writer).Encode(responses)

}
