package response

import "time"

type ResultResponse struct {
	DrawSystemId   int       `json:"drawSystemId"`
	DrawDate       time.Time `json:"drawDate"`
	GameType       string    `json:"gameType"`
	Results        []int     `json:"results"`
	SpecialResults []int     `json:"specialResults,omitempty"`
}
