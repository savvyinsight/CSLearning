package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	ID     int
	Gender string
	Name   string
}

 
type Class struct {
	Title    string
	Students []Student
}

func main() {
	c := &Class{
		Title:    "001",
		Students: make([]Student, 0, 200),
	}

	for i := 0; i < 10; i++ {
		stu := Student{
			Name:   fmt.Sprintf("stu%02d", i),
			Gender: "男",
			ID:     i,
		}
		c.Students = append(c.Students, stu)

	} //JSON 序列化：结构体-->JSON 格式的字符串
	data, err := json.Marshal(c)
	if err != nil {
		fmt.Println("json marshal failed")
		return
	}
	fmt.Printf("json:%s\n", data)
}
