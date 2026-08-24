package results

import (
	"basic_api/client"
	"basic_api/model"
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
