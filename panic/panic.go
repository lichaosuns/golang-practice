package main

import (
	"fmt"
)

func main() {

	var name string

	defer func() {
		fmt.Println("defer function 1")
	}()

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("erro occured")
		}
	}()

	if name == "" {
		panic("name is empty")
	}
}
