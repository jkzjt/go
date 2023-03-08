package model

import "time"

type Course struct {
	ID        uint
	Title     string
	Period    uint16
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (Course) TableName() string {
	return "course_test"
}
