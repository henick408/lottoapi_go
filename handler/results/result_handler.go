package results

import (
	"encoding/json"
	"lottoapi/model"
	"lottoapi/response"
	"lottoapi/service/results"
	"lottoapi/util"
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

	if results == nil {
		http.Error(writer, "No results found for given date", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	responses := mapToResponses(results)

	json.NewEncoder(writer).Encode(responses)

}

func (handler *ResultHandler) GetByGameHandler(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	gameTypeStr := request.PathValue("gameType")

	gameType, err := model.GameTypeFrom(gameTypeStr)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	results, err := handler.service.GetResulstByGame(gameType)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	responses := mapToResponses(results)
	json.NewEncoder(writer).Encode(responses)
}

func mapToResponses(results []model.DrawResult) (responses []response.ResultResponse) {
	for _, result := range results {
		responses = append(responses, result.ToResponse())
	}
	return
}
