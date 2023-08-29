package main

import (
	"fmt"
	"reflect"
)

func main() {
	variadicFunction("hello", "golang", "variadict")
	variadicFunction()
	fmt.Println("no arguments")

	variadicFunction1("1", 1, 2, 3)
	variadicFunction("2")
	fmt.Println("-----------")
	variadicFunction3("hello", 10, map[string]int{"hello": 1}, true)
}

func variadicFunction(s ...string) {
	fmt.Println("===========")
	for _index := range s {
		fmt.Println(string(s[_index]))
	}
	fmt.Println(s)
	fmt.Println("===========")
}

func variadicFunction1(key string, value ...int) {

	for _, val := range value {
		fmt.Println(key, val)
	}
}

func variadicFunction3(value ...interface{}) {
	for _, v := range value {
		fmt.Println(v, reflect.TypeOf(v), reflect.TypeOf(v).Kind())
	}
}
