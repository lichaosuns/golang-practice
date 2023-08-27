package main

import (
	"fmt"
	"strconv"
)

func main() {

	flotStr1 := "1.111111111"
	floatVal1, err := strconv.ParseFloat(flotStr1, 32)
	fmt.Println(floatVal1, err)
	fmt.Printf("%T\n", floatVal1)

}
