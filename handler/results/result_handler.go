package results

import (
	"basic_api/response"
	"basic_api/service/results"
	"encoding/json"
	"net/http"
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

	results, err := handler.service.GetLastResults()
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
	}

	writer.Header().Set("Content-Type", "application/json")

	var responses []response.ResultResponse

	for _, result := range results {
		responses = append(responses, result.ToResponse())
	}

	json.NewEncoder(writer).Encode(responses)

}
