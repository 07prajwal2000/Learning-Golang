package main

import "fmt"

func VariablesTest() {
	var name string // if we don't want to initialize the variable
	age := 22       // Initializing inline, this will auto-detect the type
	var (
		address string
		ph_no   string
	) // defining multiple variables

	num1, num2, str1 := 2, 3, "STRING"               // defining and initializing multiple variables
	const constant_number int32 = 55                 // defining constant number
	const constant_string string = "CONSTANT STRING" // defining constant string

	fmt.Println(name, age)
	fmt.Println(address, ph_no)
	fmt.Println(num1, num2, str1)
	fmt.Println(constant_number)
	fmt.Println(constant_string)
}
