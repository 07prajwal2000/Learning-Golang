package main

import "fmt"

func mapTest() map[string]int32 {
	mapp := map[string]int32{
		"Hello":  2,
		"Hello1": 3,
	}
	return mapp
}

func MapTest() {
	mapTest := mapTest()
	for k, v := range mapTest {
		fmt.Println("KEY - ", k, " , VALUE - ", v)
	}
}
