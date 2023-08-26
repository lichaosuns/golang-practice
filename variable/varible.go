package main

import (
	"fmt"
	"reflect"
)

var global_var int
var global_str string

var (
	global_a int     = 10
	global_b string  = "hello"
	global_c float32 = 3.1
	global_d string  = "unused"
)

var bb = 10

func main() {
	// 直接初始化 语法：var name type = value
	var var1 int = 10
	var var2 int = 20
	fmt.Print(var1)
	fmt.Print((var2))

	// 间接初始化

	var var3 int
	var var4 int

	var3 = 100
	var4 = 10
	fmt.Print(var3)
	fmt.Print(var4)

	// 缺省类型
	var var5 = 10
	var var6 = "hello"

	fmt.Print(reflect.TypeOf(var5))
	fmt.Print(reflect.TypeOf(var6))

	// 一行多个

	var a, b, c string = "hello", "golang", "bye"

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	m, n, z := 1, 2, "h"
	fmt.Println(m)
	fmt.Println(n)
	fmt.Println(z)

	//全局变量
	fmt.Println(global_var)
	fmt.Println(global_str)

	local_var := 10
	// 局部变量
	fmt.Println(local_var)

	fmt.Println(global_a)
	fmt.Println(global_c)
	fmt.Println(global_b)

}
