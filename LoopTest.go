package main

import "fmt"

func LoopTest() {
	// For loop
	fmt.Println("For loop")
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	fmt.Println("----------")
	fmt.Println("Foreach loop like c#")
	// Foreach Loop like in c#
	arr := []int{1, 2, 3, 4}
	for _, num := range arr {
		fmt.Println(num)
	}

	fmt.Println("----------")
	fmt.Println("While loop")
	i := 0
	for {
		if i > 5 {
			break
		}
		fmt.Println(i)
		i++
	}

	fmt.Println("----------")
	fmt.Println("Switch Statement")
	switch i {
	case 5:
		fmt.Println("Five")

	case 6:
		fmt.Println("Six")

	default:
		fmt.Println("Other number")
	}

	fmt.Println("___________")
	fmt.Println("Switch Statement without arg")
	switch {
	case i == 5:
		fmt.Println("Five")

	case i > 5:
		fmt.Println("Greater than five")

	default:
		fmt.Println("Other number")
	}

}
