package models

import (
	"gorm.io/gorm"
)

type Result struct {
	gorm.Model
	Id          int
	AthleteId   int
	Ordinal     int
	Score       int
	Description string
}

type Athlete struct {
	gorm.Model
	Id           int
	CompetitorId int
	Name         string
	Age          int
	Gender       string
}
