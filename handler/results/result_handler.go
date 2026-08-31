package results

import (
	"lottoapi/model"
	"lottoapi/response"
	"lottoapi/service/results"
	"lottoapi/util"
	"net/http"
	"time"

	"github.com/labstack/echo/v5"
)

type ResultHandler struct {
	service *results.ResultService
}

func NewResultHandler(service *results.ResultService) *ResultHandler {
	return &ResultHandler{
		service: service,
	}
}

func (handler *ResultHandler) GetAllHandler(context *echo.Context) error {
	drawDateStr := context.QueryParam("drawDate")
	var (
		results []model.DrawResult
		err     error
	)

	if drawDateStr != "" {
		date, parseErr := time.Parse(time.DateOnly, drawDateStr)
		if parseErr != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Incorrect date.")
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
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if results == nil {
		return echo.NewHTTPError(http.StatusNotFound, "No result found for given date")
	}

	responses := mapToResponses(results)

	return context.JSON(http.StatusOK, responses)

}

func (handler *ResultHandler) GetByGameHandler(context *echo.Context) error {
	gameTypeStr := context.Param("gameType")
	gameType, gameTypeErr := model.GameTypeFrom(gameTypeStr)
	if gameTypeErr != nil {
		return echo.NewHTTPError(http.StatusNotFound, gameTypeErr.Error())
	}

	var (
		result model.DrawResult
		err    error
	)

	drawDateStr := context.QueryParam("drawDate")

	if drawDateStr != "" {
		date, parseErr := time.Parse(time.DateOnly, drawDateStr)
		if parseErr != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Incorrect date.")
		}
		drawDate := util.NewDate(
			date.Year(),
			int(date.Month()),
			date.Day(),
		)
		result, err = handler.service.GetResultByDateByGame(drawDate, gameType)
	} else {
		result, err = handler.service.GetResultByGame(gameType)
		if result.DrawSystemId == 0 {
			return echo.NewHTTPError(http.StatusNotFound, "No result found for given date")
		}
	}

	if result.DrawSystemId == 0 {
		return echo.NewHTTPError(http.StatusNotFound, "No result found")
	}

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	response := result.ToResponse()
	return context.JSON(http.StatusOK, response)
}

func mapToResponses(results []model.DrawResult) (responses []response.ResultResponse) {
	for _, result := range results {
		responses = append(responses, result.ToResponse())
	}
	return
}
