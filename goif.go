package main

import "fmt"

func main() {
	//if
	//if:true false
	if false {
		fmt.Println("yes")
	} else {
		fmt.Println("no")
	}

	//if yes no
	var pop int
	fmt.Print("你要多少錢:")
	fmt.Scanln(&pop)
	if pop < 300 {
		fmt.Println("太少")
	} else if pop <= 1000 {
		fmt.Println("yes")
	} else {
		fmt.Println("太多")
	}
	fmt.Println("結束")
}
