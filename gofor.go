package main

import "fmt"

func main() {
	// var io int
	// fmt.Print("輸入:")
	// fmt.Scanln(&io)
	// for io >= 5 {
	// 	if io > 3 {
	// 		fmt.Println("halo")
	// 	} else {
	// 		fmt.Println("0")
	// 	}
	// }

	//print 0~2
	var pop int = 0
	for pop < 3 {
		fmt.Println(pop)
		pop++
	}
	//print奇數
	var pop1 int
	for pop1 = 1; pop1 < 50; pop1 += 2 {
		fmt.Println(pop1)
	}

	//1+~50
	var pop2 int = 0
	var x int = 1
	for x <= 50 {
		pop2 += x
		x++
	}
	fmt.Println(pop2)

	//break
	var pop3 int = 0
	// fmt.Print("輸入:")
	// fmt.Scanln(&pop3)
	for pop3 < 8 {
		//0~7 偵測到4就停止
		if pop3 == 4 {
			break
		}
		fmt.Println(pop3)
		pop3++
	}
	//countinue
	var pop4 int
	for pop4 = 0; pop4 < 8; pop4++ {
		//0~7 略過5 繼續執行
		if pop4 == 5 {
			continue
		}
		fmt.Println(pop4)
	}
	//print偶數
	var pop5 int
	for pop5 = 0; pop5 < 50; pop5++ {
		if pop5%2 == 0 {
			fmt.Println(pop5)
		}

	}
	//print奇數
	var pop6 int
	for pop6 = 0; pop6 < 50; pop6++ {
		if pop6%2 == 0 { //pop6是偶數,不印出來
			continue
		}
		fmt.Println(pop6)
	}
	//輸入特定數 停止stop
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
