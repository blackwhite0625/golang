package main

import "fmt"

// 印出"hello"副程式 halo函式
func halo() {
	println("hello")
}

// 印出任何訊息函式
// 參數名稱 資料型態
func haprint(msg string) {
	fmt.Println(msg)
}

// 定義兩數相加
func add(num1 int, num2 int) {
	var ans int = num1 + num2
	fmt.Println(ans)
}
func sum() {
	var popo int = 0
	for true { // 無窮迴圈
		var pop7 int
		fmt.Print("請輸入:")
		fmt.Scanln(&pop7)
		if pop7 == 4 {
			break
		}
		//加總
		popo += pop7
	}
	fmt.Println(popo)
}
func mux(max int) {
	var sum1 int = 0
	var x int
	for x = 1; x <= max; x++ {
		sum1 += x
	}
	fmt.Println(sum1)
}
func main() {
	//呼叫halo
	halo()
	//呼叫haprint
	haprint("hello")
	//呼叫add
	add(5, 2)
	sum()
	mux(10)//1+2...+10
}
