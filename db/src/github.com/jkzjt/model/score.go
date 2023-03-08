package model

import "time"

type Score struct {
	ID        uint
	No        string
	CourseId  uint
	Point     float64
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (Score) TableName() string {
	return "score_test"
}
