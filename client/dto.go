package client

import (
	"basic_api/model"
	"time"
)

type drawDto struct {
	DrawSystemId         int64
	DrawDate             time.Time
	GameType             string
	multiplierValue      int
	Results              []gameResultDto
	showSpecialResults   bool
	isNewEuroJackpotDraw bool
}

type gameResultDto struct {
	DrawDate       time.Time
	DrawSystemId   int64
	GameType       string
	ResultsJson    []int
	SpecialResults []int
}

func (drawDto drawDto) toModel() model.DrawResult {
	var resultDto gameResultDto
	for _, gameDto := range drawDto.Results {
		if gameDto.GameType == drawDto.GameType {
			resultDto = gameDto
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
