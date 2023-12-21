package main

import (
	"fmt"
)

// Student 結構表示一個學生的基本信息
type Student struct {
	ID   int
	Name string
	Age  int
}

// 學生資料庫，用 slice 來保存學生
var students []Student

// 新增學生到資料庫
func Add(id int, name string, age int) {
	student := Student{
		ID:   id,
		Name: name,
		Age:  age,
	}
	students = append(students, student)
	fmt.Println("學生已新增:", student)
}

// PrintStudents 函式用於列印所有學生的資訊
func PrintStudents() {
	fmt.Println("所有學生資訊:")
	for _, student := range students {
		fmt.Printf("座號: %d, 姓名: %s, 年齡: %d\n", student.ID, student.Name, student.Age)
	}
}

func main() {
	// 新增一些範例學生
	Add(1, "黑白", 18)
	Add(2, "珍珠奶茶", 19)
	Add(3, "iphone", 20)

	// 列印所有學生資訊
	PrintStudents()
}
