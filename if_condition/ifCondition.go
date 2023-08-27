package main

import (
	"fmt"
)

func main() {

	var str = "hello"

	x := true

	if x == true {
		fmt.Println(str)
	} else {
		fmt.Println("else")
	}

	if true {
		fmt.Println("always true")
	}

	// if 语句初始化

	if newVal := 10; true {
		fmt.Println(newVal)
	}
}
