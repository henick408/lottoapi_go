package model

import "time"

type DrawResult struct {
	//ID             int64
	DrawSystemId   int64
	DrawDate       time.Time
	GameType       GameType
	Results        []int
	SpecialResults []int
}
