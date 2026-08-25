package client

import (
	"lottoapi/model"
	"sort"
	"time"
)

type DrawDto struct {
	DrawSystemId         int64
	DrawDate             time.Time
	GameType             string
	multiplierValue      int
	Results              []GameResultDto
	showSpecialResults   bool
	isNewEuroJackpotDraw bool
}

type GameResultDto struct {
	DrawDate       time.Time
	DrawSystemId   int64
	GameType       string
	ResultsJson    []int
	SpecialResults []int
}

func (drawDto DrawDto) ToModel() model.DrawResult {
	var resultDto GameResultDto
	for _, gameDto := range drawDto.Results {
		if gameDto.GameType == drawDto.GameType {
			resultDto = gameDto
			sort.Ints(resultDto.ResultsJson)
			sort.Ints(resultDto.SpecialResults)
			break
		}
	}

	return model.DrawResult{
		DrawSystemId:   drawDto.DrawSystemId,
		DrawDate:       drawDto.DrawDate,
		GameType:       model.GameType(drawDto.GameType),
		Results:        resultDto.ResultsJson,
		SpecialResults: resultDto.SpecialResults,
	}
}
