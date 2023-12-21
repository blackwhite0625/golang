package main

//輸出輸入函式庫
import "fmt"

// main函式 程式進入點
func main() {
	//資料型態data
	fmt.Println(5)       //int
	fmt.Println(3.14159) //float64
	fmt.Println("字串")    //string字串
	fmt.Println(true)    //bool布林
	fmt.Println('a')     //rune字符

	//變數var
	//整數int
	var pop int // 宣告變數
	pop = 5
	fmt.Println(pop)
	//變數可被替換
	pop = 89
	fmt.Println(pop)

	//浮點數float64
	var pop1 float64
	pop1 = 3.14159
	fmt.Println(pop1)
	//資料型態=適當的資料data
	var pop2 float64 = 84.32
	fmt.Println(pop2)

	//字串string
	var pop3 string = "你好lolpo"
	fmt.Println(pop3)

	//布林值bool
	var pop4 bool = true
	fmt.Println(pop4)

	//字符rune
	var pop5 rune = 'g'
	fmt.Println(pop5)

}
