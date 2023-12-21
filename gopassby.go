package main

import "fmt"

// 參數傳遞
// 接收參數pop,進行+10
// pass by value
func add(pop int) {
	//pop = pop + 10
	pop += 10
	fmt.Println("valueadd:", pop)
}

// 指標參數傳遞
// 接收指標參數ptr
// 反解後進行+10
// pass by pointer
func add2(ptr *int) {
	//*xptr=*xptr+10
	*ptr += 10
	//fmt.Println("pointeradd:", *ptr)
}

func main() {
	// 呼叫函式,傳遞參數
	// pass by value
	var x int = 10
	add(x)
	fmt.Println("valuemain:", x)

	//呼叫函式,建立並傳入指標參數ptr
	//pass by pointer
	var pop1 int = 10
	// var ptr *int = &pop1
	// add2(ptr)
	add2(&pop1)
	fmt.Println("pointermain:", pop1)

	//使用者輸入使用到指標參數pass by pointer
	var test string
	var testptr *string = &test
	fmt.Print("輸入:")
	// fmt.Scanln(&test)
	fmt.Scanln(testptr)
	fmt.Println(test)

}
