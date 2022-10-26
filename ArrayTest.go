package main

import "fmt"

func arrayTest() []int {
	arr := []int{1, 2, 3, 4}
	return arr
}

func ArrayTest() {
	arr := arrayTest()
	fmt.Println(arr)
	a := arr[:2]
	a[0] = 4
	fmt.Println(arr)

	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}

	for i, num := range arr {
		fmt.Printf("i - %d, Number %d\n", i, num)
	}

	fmt.Println(fmt.Sprint(22)) // Converts 22 int to "22" string -> fmt.Sprint(22)
}
