package main

import (
	"fmt"
)

func main() {
	// 延迟执行
	defer second()

	first()
	x := 0
	// 这种场景下就不会执行
	// os.Exit(0)

	fmt.Println(0 / 3 / x)
}

func first() {
	fmt.Println("first function")
}

func second() {
	fmt.Println("second function")
}
