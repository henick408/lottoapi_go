package client

import (
	"lottoapi/model"
	"lottoapi/util"
	"os"

	"github.com/go-resty/resty/v2"
)

var baseURL = "https://developers.lotto.pl/api/open/v1/lotteries"

func getLottoSecret() string {
	return os.Getenv("LOTTO_API_KEY")
}

// type LottoClient struct {
// 	httpClient *http.Client
// 	baseURL    string
// }

type LottoClient struct {
	Client *resty.Client
}

func NewLottoClient(client *resty.Client) *LottoClient {
	client.SetBaseURL(baseURL)
	client.SetHeader("secret", getLottoSecret())
	client.SetHeader("Accept", "application/json")
	return &LottoClient{
		Client: client,
	}
}

func (lotto *LottoClient) GetLastResults() ([]model.DrawResult, error) {
	var drawDtos []DrawDto

	_, err := lotto.Client.NewRequest().
		SetResult(&drawDtos).
		Get("/draw-results/last-results")
	if err != nil {
		return nil, err
	}

	drawResults := flatMapToModels(drawDtos)
	return drawResults, nil
}

func (lotto *LottoClient) GetResultsByDate(drawDate util.Date) ([]model.DrawResult, error) {
	var drawDtos []DrawDto
	_, err := lotto.Client.NewRequest().
		SetQueryParam("drawDate", drawDate.ToString()).
		SetResult(&drawDtos).
		Get("/draw=results/by-date")
	if err != nil {
		return nil, err
	}
	drawResults := flatMapToModels(drawDtos)
	return drawResults, nil
}

func (lotto *LottoClient) GetResultByGame(gameType model.GameType) (model.DrawResult, error) {
	var drawDtos []DrawDto

	_, err := lotto.Client.NewRequest().
		SetQueryParam("gameType", string(gameType)).
		SetResult(&drawDtos).
		Get("draw-results/last-results-per-game")
	if err != nil {
		return model.DrawResult{}, err
	}
	drawDto := drawDtos[0]
	drawResult := drawDto.ToModel()

	return drawResult, nil
}
