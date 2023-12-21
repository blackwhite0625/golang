package main

import "fmt"

func main() {
	//算術運算 + - * /
	var pop int
	pop = (((6 + 2 - 1) / 2) * 6)
	fmt.Println(pop)

	//指定運算 = += -= *= /=
	pop = 23
	pop = pop + 3
	pop += 2 //pop = pop +2
	fmt.Println(pop)

	//單元運算 ++ --
	pop++ // pop + 1
	fmt.Println(pop)

	//比較運算bool > < >= <= ==
	var bool1 bool = 3 <= 5
	fmt.Println(bool1)

	//邏輯運算 ! && ||
	bool1 = !true //!反運算
	fmt.Println(bool1)
	bool1 = true && true //且and&& 兩邊都是true才是true ,不然都是false
	fmt.Println(bool1)
	bool1 = true || false //或or|| 只要一邊是true ,就是true
	fmt.Println(bool1)

	//簡短變數宣告:=
	pop1 := 5
	fmt.Println(pop1)

}
