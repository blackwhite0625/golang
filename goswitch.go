package main

import "fmt"

func main() {
	var pop int
	fmt.Print("你要多少錢:")
	fmt.Scanln(&pop)
	switch {
	case pop < 300:
		fmt.Println("太少")
	case pop <= 1000:
		fmt.Println("yes")
	default:
		fmt.Println("太多")
	}
}
