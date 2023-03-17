package models

import (
	"encoding/json"
	"strconv"
)

type FlexInt int

func (fi *FlexInt) UnmarshalJSON(b []byte) error {
	if b[0] != '"' {
		return json.Unmarshal(b, (*int)(fi))
	}
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	i, err := strconv.Atoi(s)
	if err != nil {
		return err
	}
	*fi = FlexInt(i)
	return nil
}

type Entrant struct {
	CompetitorId   FlexInt
	CompetitorName string
	Gender         string
	Age            FlexInt
}

type Score struct {
	Ordinal      FlexInt
	Rank         FlexInt
	Score        FlexInt
	ScoreDisplay string
}

type LeaderBoardRow struct {
	OverallRank  FlexInt
	OverallScore FlexInt
	Entrant      Entrant
	Scores       []Score
}

type Pagination struct {
	TotalPages       FlexInt
	TotalCompetitors FlexInt
	CurrentPage      FlexInt
}

type Response struct {
	Pagination      Pagination
	LeaderboardRows []LeaderBoardRow
}
