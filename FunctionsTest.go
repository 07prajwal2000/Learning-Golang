package main

import (
	"fmt"
)

/*
Using some system resource like socket, file system.
returns some resource id
*/
func use_resource() string {
	res_id := "sys_resource_fs"
	fmt.Println("Using Resource - ", res_id)
	return res_id
}

/*
defer keyword is like finally in c#. it will be called after the execution of the current function.
*/
func FunctionTest() error {
	resource := use_resource()
	defer clean_up(resource)
	fmt.Println("Called after defer")
	resource2 := use_resource()
	defer clean_up(resource2)
	return fmt.Errorf("this is like an exception")
}

func clean_up(str string) {
	fmt.Printf("\033[31mCleaning %s \033[0m ", str)
}
