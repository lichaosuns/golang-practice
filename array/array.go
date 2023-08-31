package main

import (
	"fmt"
	"reflect"
)

func main() {

	var stringArr [2]string

	stringArr[0] = "hello"
	stringArr[1] = "java"

	for _, v := range stringArr {
		fmt.Println(v)
	}

	fmt.Println(reflect.ValueOf(stringArr).Kind(), reflect.ValueOf(stringArr), reflect.TypeOf(stringArr))

	preInit := [5]int{1, 2, 3, 4, 5}

	fmt.Println(preInit)

	var preInit2 = [5]int{1, 2, 3, 4, 5}

	fmt.Println(preInit2)

	var partInit = [5]int{1, 2, 3}
	fmt.Println(partInit)

	// 省略号

	var preInit3 = [...]int{1, 2, 3, 4}

	fmt.Println(preInit3)

	// 复制数组

	var arrA = [2]string{"hello", "golang"}
	arrB := arrA
	arrC := &arrA

	fmt.Println(arrA, arrB, arrC)

	arrB[0] = "no Hello"
	arrC[1] = "java"

	fmt.Println(arrA, arrB, arrC)
	fmt.Println(&arrA == arrC)
	fmt.Println(&arrA == &arrB)

	flag := itemExits([5]int{1, 2, 4, 5, 6}, 1)
	fmt.Println(flag)

	// 过滤数组

	contries := [4]string{"china", "england", "aaa"}

	fmt.Println(contries[:1])
	fmt.Println(contries[:2])
	fmt.Println(contries[1:])
	fmt.Println(contries[0:0])
}

func itemExits(arrV interface{}, item interface{}) bool {
	arrVType := reflect.ValueOf(arrV)
	if arrVType.Kind() != reflect.Array {
		panic("invalid type")
	}

	fmt.Println(arrVType.Kind())

	for i := 0; i < arrVType.Len(); i++ {
		if arrVType.Index(i).Interface() == item {
			return true
		}
	}
	return false
}
