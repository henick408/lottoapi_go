package client

import (
	"basic_api/model"
	"basic_api/util"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
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

func (client *LottoClient) get(url string) (resp *http.Response, err error) {
	request, err := http.NewRequest(http.MethodGet, url, nil)

	if err != nil {
		return nil, err
	}
	return client.do(request)
}
func (client *LottoClient) do(request *http.Request) (*http.Response, error) {
	request.Header.Set("secret", getLottoSecret())
	request.Header.Set("accept", "application/json")
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

	var drawDtos []DrawDto

	err = json.NewDecoder(response.Body).Decode(&drawDtos)
	if err != nil {
		return nil, err
	}

	drawResults := mapDrawDtos(drawDtos)

	return drawResults, nil

}

func (client *LottoClient) GetResultsByDate(drawDate util.Date) ([]model.DrawResult, error) {
	params := url.Values{}
	params.Set("drawDate", drawDate.ToString())
	url := fmt.Sprintf(
		"%s/draw-results/by-date?%s",
		baseURL, params.Encode(),
	)
	response, err := client.get(url)
	if err != nil {
		return nil, err
	}

	var drawDtos []DrawDto

	err = json.NewDecoder(response.Body).Decode(&drawDtos)
	if err != nil {
		return nil, err
	}

	drawResults := mapDrawDtos(drawDtos)

	return drawResults, nil
}

func mapDrawDtos(drawDtos []DrawDto) (drawResults []model.DrawResult) {
	for _, drawDto := range drawDtos {
		drawResults = append(drawResults, drawDto.ToModel())
	}
	return
}
