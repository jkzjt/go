package mapper

import (
	"github.com/jkzjt/mygo/db/model"
)

func FindAll() (students []model.Student) {
	return FindByConds()
}

func FindByConds(conds ...any) (students []model.Student) {
	result := DB.Find(&students, conds...)
	if result.Error != nil {
		panic("table[student_test]: Find fail")
	}
	return
}
