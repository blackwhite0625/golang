package main

import "fmt"

func main() {
	//建立data變數
	var x int = 5

	//存放到指標變數,指標資料型態 *資料型態
	var y *int = &x

	//反解指標變數,反解運算 *指標變數名稱
	var z *int = &x

	//取得記憶體位置:&變數名稱
	fmt.Println(&x)
	fmt.Println(&y)
	//反解運算 *指標變數名稱
	fmt.Println(*z)

	var test string = "halo"
	fmt.Println(test)
	var testptr *string = &test
	fmt.Println(testptr)
	//fmt.Println(*testptr)
	test = *testptr
	fmt.Println(test)
}
