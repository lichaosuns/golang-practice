package main

import (
	"fmt"
	"strconv"
)

func main() {

	boolStr := "true"
	boolVal, err := strconv.ParseBool(boolStr)
	fmt.Println(boolVal)
	fmt.Println(err)

	boolStr1 := "t"
	boolVal1, err := strconv.ParseBool(boolStr1)

	fmt.Println(boolVal1)

	boolStr2 := "T"
	boolVal2, err := strconv.ParseBool(boolStr2)

	fmt.Println(boolVal2)

	boolFalse := "0"
	boolFalseVal, err := strconv.ParseBool(boolFalse)
	fmt.Println(boolFalseVal)

	var bitA = 0b111
	var bitb = 0b011

	fmt.Println(bitA & bitb)

}
