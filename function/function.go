package main

import (
	"fmt"
)

func main() {
	fmt.Println("this is main function")
	SimpleFunction()
	add(1, 10)

	result := addAndReturn(12, 20)
	fmt.Println("result is ", result)

	result = addAndReturnNoName(12, 20)
	fmt.Println("result is ", result)

	//functionThrow()

	a := 10
	b := 11
	functionNoAddress(a, b)

	fmt.Println("no address", a, b)

	functionAddress(&a, &b)
	fmt.Println("address", a, b)

	var noNameFun = func(x int, y int) (int, string) {
		fmt.Println("i am a no name")
		return x + y, "hello"
	}

	fmt.Println(noNameFun(10, 20))

	func(x int, y int) {
		fmt.Println(x, "---", y)
	}(10, 20)

	var outSideVar = 10

	// 闭包
	func() {
		fmt.Println("clousure:", outSideVar)
	}()

	// high order function

	function := returnFunc(10)
	fmt.Println("function ", function(1))

	var twoFuncRes = twoFunc(1)(2)(3)
	fmt.Println((twoFuncRes))

}

func SimpleFunction() {
	fmt.Println("this is a simple function")
}

func add(x int, y int) {
	fmt.Println("add result:", x+y)
}

func addAndReturn(x int, y int) (z int) {
	return x + y
}

func functionThrow() {
	panic("erro")
}

func addAndReturnNoName(x int, y int) int {
	return x + y
}

func functionNoAddress(x int, y int) {
	x = x + 10
	y = y + 11
}

func functionAddress(x *int, y *int) {
	*x = *x + 10
	*y = *y + 11
	// x表示的是入参指针的地址值
	// &x 表示的是自身的地址值
	fmt.Println(x, "=", &x)
	fmt.Println(*x)
	fmt.Println("---------")
}

func returnFunc(x int) func(int) int {
	return func(y int) int {
		return x + y
	}
}

func twoFunc(z int) func(int) func(int) int {
	return func(x int) func(int) int {
		return func(y int) int {
			return x + y + z
		}
	}
}
