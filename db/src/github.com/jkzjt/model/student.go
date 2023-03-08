package model

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"
)

type Student struct {
	ID        uint      `json:"id"`
	IDNo      string    `gorm:"column:id_no" json:"id_no"`
	No        string    `json:"no"`
	Name      string    `json:"name"`
	Age       uint8     `json:"age"`
	Gender    string    `json:"gender"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (Student) TableName() string {
	return "student_test"
}

func (student *Student) String() string {
	b, err := json.Marshal(*student)
	if err != nil {
		return fmt.Sprintf("%+v", *student)
	}
	var out bytes.Buffer
	err = json.Indent(&out, b, "", " ")
	if err != nil {
		return fmt.Sprintf("%+v", *student)
	}
	return out.String()
}
