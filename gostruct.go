package main

import "fmt"

// 結構定義
type Test struct {
	x int
	y int
}
type people struct {
	id   int
	name string
	age  int
}

func main() {
	//Test struct
	// var p1 Test = Test{3, 4}
	// fmt.Println(p1.x, p1.y)
	var p1 Test = Test{3, 4}
	var p2 Test = Test{y: 2, x: 3}
	fmt.Println(p1.x, p1.y)
	fmt.Println(p1)
	fmt.Println(p2.x, p2.y)

	//people struct
	var pop people = people{29, "黑白", 18}
	fmt.Println(pop.id, pop.name, pop.age)
	var pop1 people = people{id: 29, name: "黑白", age: 18}
	pop1.name = "阿比"
	fmt.Println(pop1.id, pop1.name, pop1.age)

}
