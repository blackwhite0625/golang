package main

import "fmt"

// 結束函式stop func
func hey(uin string) {
	// fmt.Print("輸入:")
	// fmt.Scanln(&uin)

	if uin == "你好" {
		fmt.Println("對")
	} else {
		fmt.Println("錯")
		//執行return 結束程式
		return
	}
	fmt.Println(uin)
}

// 函式回傳值(單一)
func hal(n1 int, n2 int) int {
	var n3 int = n1 * n2
	//fmt.Println(n3)
	//n3 := n1 * n2 (:=簡短var xx int = xx)
	return n3
}

// 函式回傳值(多個)
func funi() (int, string) {
	return 10, "data"

}
func main() {
	//結束函式stop func
	hey("halo")
	hey("你好")

	//函式回傳值(單一)
	var n4 int = hal(4, 5)
	fmt.Println(n4 + 44)
	//n4 := hal(4, 5)

	//函式回傳值(多個)
	var intt int
	var stringg string
	intt, stringg = funi()
	fmt.Println(intt, stringg)

}
