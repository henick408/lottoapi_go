package model

import (
	"fmt"
	"strings"
)

type GameType string

const (
	GameTypeLotto       GameType = "Lotto"
	GameTypeLottoPlus   GameType = "LottoPlus"
	GameTypeMiniLotto   GameType = "MiniLotto"
	GameTypeEuroJackpot GameType = "EuroJackpot"
)

func GameTypeFrom(value string) (GameType, error) {
	value = strings.ToLower(value)
	switch value {
	case "lotto":
		return GameTypeLotto, nil
	case "minilotto":
		return GameTypeMiniLotto, nil
	case "lottoplus":
		return GameTypeLottoPlus, nil
	case "eurojackpot":
		return GameTypeEuroJackpot, nil
	default:
		return "", fmt.Errorf("Unknown game type: %q", value)
	}
}
