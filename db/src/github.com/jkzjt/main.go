package main

import (
	"fmt"

	"github.com/jkzjt/mygo/db/mapper"
	"github.com/jkzjt/mygo/db/model"
)

func main() {
	students := mapper.FindAll()
	fmt.Println(students)

	students = mapper.FindByConds(model.Student{Gender: "1"})
	fmt.Println(students)

	students = mapper.FindByConds(map[string]any{"Gender": "0", "Age": 18})
	fmt.Println(students)
}
