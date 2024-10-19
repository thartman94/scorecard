package models

import (
	"errors"
	"fmt"
)

type Team struct {
	name     string
	location string
	abv      string
}

func (team Team) LongName() string {
	return fmt.Sprintf("%s %s", team.location, team.name)
}

type Game struct {
	homeTeam Team
	awayTeam Team
	// data
	// startTime
	// finishTime
	// weather
}

type Inning struct {
	game Game
}

type InningHalf struct {
	inning Inning
}

type Player struct {
	FirstName string
	LastName  string
	Number    int
	Team      Team
}

type PlateAppearance struct {
	inningHalf InningHalf
	batter     Player
	strikes    int
	balls      int
}

func (pa *PlateAppearance) AddBall() (*int, error) {
	if pa.balls == 4 {
		return nil, errors.New("bro has already walked")
	}

	pa.balls++
	return &pa.balls, nil
}

func (pa *PlateAppearance) MinusBall() (*int, error) {
	if pa.balls == 0 {
		return nil, errors.New("0 is the smallest number")
	}

	pa.balls--
	return &pa.balls, nil
}

func (pa *PlateAppearance) AddStrike() (*int, error) {
	if pa.strikes == 3 {
		return nil, errors.New("bro has already struck out")
	}

	pa.strikes++

	return &pa.strikes, nil
}

func (pa *PlateAppearance) MinusStrike() (*int, error) {
	if pa.strikes == 0 {
		return nil, errors.New("0 is the smallest number")
	}

	pa.strikes--
	return &pa.strikes, nil
}
