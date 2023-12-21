package main

import "fmt"

func main() {
	//輸入輸出 input output
	//(資料,資料,.....)
	// fmt.Println(43, "你好", true)
	// var pop int
	// fmt.Print("輸入:")
	// fmt.Scanln(&pop)
	// fmt.Println(pop)

	var pop1 int
	var pop2 int
	fmt.Print("第一數:")
	fmt.Scanln(&pop1)
	fmt.Print("第二數:")
	fmt.Scanln(&pop2)
	var pop1and2 int = pop1 + pop2
	fmt.Println(pop1and2)

	var pop3 int
	var pop4 int
	fmt.Print("輸入兩數(空格開):")
	fmt.Scanln(&pop3, &pop4)
	var pop5 int = pop3 + pop4
	fmt.Println(pop5)

}
