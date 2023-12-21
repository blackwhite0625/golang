package main

import "fmt"

func main() {
	//陣列資料的編號
	//給定初始陣列資料
	//整數陣列
	var arr [5]int = [5]int{22, 43, 11, 565, 13}
	//字串陣列
	var test [3]string = [3]string{"python", "golang", "C"}

	//逐一給定陣列資料
	// var arr [5]int
	// arr[0] = 22
	// arr[1] = 43
	// arr[2] = 11
	// arr[3] = 565
	// arr[4] = 13
	// arr[5]=xx 不能超過陣列的長度

	//逐一取得陣列資料
	var data int
	var add int = 0
	for data = 0; data < len(arr); data++ {
		fmt.Println("巡迴陣列:", arr[data])
		add += arr[data]
	}
	fmt.Println("陣列for全加:", add)
	fmt.Println("陣列全加除len長度:", add/len(arr))
	var sum int = arr[0] + arr[1] + arr[2] + arr[3] + arr[4]
	fmt.Println("陣列依序全加:", sum)

	fmt.Println("陣列乘322:", arr[0]*322)
	//印出 陣列[編號]
	fmt.Println("陣列編號1:", arr[2])
	fmt.Println("陣列編號2:", test[1])
	//len 取得陣列長度
	fmt.Println("整數陣列長度:", len(arr))
	fmt.Println("字串陣列長度:", len(test))

	//逐一輸入陣列資料
	fmt.Println("逐一輸入:")
	var sum1 [5]int
	var index int
	for index = 0; index < len(sum1); index++ {
		fmt.Scanln(&sum1[index])
	}
	fmt.Println(sum1)
}
