package client

import (
	"basic_api/model"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

var baseURL = "https://developers.lotto.pl/api/open/v1/lotteries"

func getLottoSecret() string {
	return os.Getenv("LOTTO_API_KEY")
}

type LottoClient struct {
	httpClient *http.Client
	baseURL    string
}

func NewLottoClient(httpClient *http.Client) *LottoClient {
	return &LottoClient{
		httpClient: httpClient,
		baseURL:    baseURL,
	}
}

// func (client *LottoClient) GetLastResultsByGame(gameTypeRaw string) {
// 	gameType, err := model.GameTypeFrom(gameTypeRaw)
// 	if err != nil {
// 		return
// 	}

// 	url := fmt.Sprintf(
// 		"%s/results"
// 	)

// }

func (client *LottoClient) get(url string) (resp *http.Response, err error) {
	request, err := http.NewRequest(http.MethodGet, url, nil)

	if err != nil {
		return nil, err
	}
	return client.do(request)
}
func (client *LottoClient) do(request *http.Request) (*http.Response, error) {
	request.Header.Add("secret", getLottoSecret())
	request.Header.Add("accept", "application/json")
	return client.httpClient.Do(request)
}

func (client *LottoClient) GetLastResults() ([]model.DrawResult, error) {
	url := fmt.Sprintf(
		"%s/draw-results/last-results",
		baseURL,
	)
	response, err := client.get(url)
	if err != nil {
		return nil, err
	}

	var drawDtos []drawDto

	err = json.NewDecoder(response.Body).Decode(&drawDtos)
	if err != nil {
		return nil, err
	}

	var drawResults []model.DrawResult

	for _, drawDto := range drawDtos {
		drawResults = append(drawResults, drawDto.toModel())
	}

	return drawResults, nil

}
