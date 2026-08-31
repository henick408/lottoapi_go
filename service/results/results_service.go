package results

import (
	"lottoapi/client"
	"lottoapi/model"
	"lottoapi/util"
)

type ResultService struct {
	lottoClient *client.LottoClient
}

func NewResultService(lottoClient *client.LottoClient) *ResultService {
	return &ResultService{
		lottoClient: lottoClient,
	}
}

func (service *ResultService) GetLastResults() ([]model.DrawResult, error) {
	drawResults, err := service.lottoClient.GetLastResults()
	if err != nil {
		return nil, err
	}

	return drawResults, nil
}

func (service *ResultService) GetResultsByDate(drawDate util.Date) ([]model.DrawResult, error) {
	drawResults, err := service.lottoClient.GetResultsByDate(drawDate)
	if err != nil {
		return nil, err
	}

	return drawResults, nil
}

func (service *ResultService) GetResultByGame(gameType model.GameType) (model.DrawResult, error) {
	drawResult, err := service.lottoClient.GetLastResultByGame(gameType)
	if err != nil {
		return model.DrawResult{}, err
	}
	return drawResult, nil
}

func (service *ResultService) GetResultByDateByGame(drawDate util.Date, gameType model.GameType) (model.DrawResult, error) {
	drawResult, err := service.lottoClient.GetResultByDateByGame(drawDate, gameType)
	if err != nil {
		return model.DrawResult{}, err
	}
	return drawResult, nil
}
