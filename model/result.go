package model

import (
	"lottoapi/response"
	"time"
)

type DrawResult struct {
	//ID             int64
	DrawSystemId   int64
	DrawDate       time.Time
	GameType       GameType
	Results        []int
	SpecialResults []int
}

func (draw DrawResult) ToResponse() response.ResultResponse {
	return response.ResultResponse{
		DrawSystemId:   int(draw.DrawSystemId),
		DrawDate:       draw.DrawDate,
		GameType:       string(draw.GameType),
		Results:        draw.Results,
		SpecialResults: draw.SpecialResults,
	}
}
