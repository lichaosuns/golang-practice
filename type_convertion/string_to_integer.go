package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	intStr := "100"

	// 自动识别
	intVal, err := strconv.Atoi(intStr)
	fmt.Println(intVal, err, reflect.TypeOf(intVal))

	intStr = "1aa"
	// 自动识别
	intVal, err = strconv.Atoi(intStr)
	fmt.Println(intVal, err, reflect.TypeOf(err), reflect.TypeOf(intVal))

	// 转换

	intStr2 := "0b1000"

	int64Val, err := strconv.ParseInt(intStr2, 0, 8)

	fmt.Println(int64Val, err, reflect.TypeOf(int64Val))

	// 从输入中获取

	intVal10 := 10

	returnRes, err1 := fmt.Sscan("1000", &intVal10)

	fmt.Println(intVal10)
	fmt.Println(err1)
	fmt.Println(returnRes)

}
