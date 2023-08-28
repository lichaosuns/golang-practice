package main

import (
	"fmt"
)

func main() {
	for k := 0; k <= 10; k++ {
		fmt.Println(k)
	}

	var k = 0

	for k <= 10 {
		fmt.Println("********", k, "*****")
		k++
	}

	strDict := map[string]string{"hello": "world", "golang": "good"}

	for _key, _value := range strDict {
		fmt.Println(_key, ":", _value)
	}

	fmt.Println(strDict["hello"])

	var iteratorStr = "hello"

	for _index := range iteratorStr {
		fmt.Printf("%c\n", iteratorStr[_index])
		fmt.Println(string(iteratorStr[_index]), " is encode")
	}
}
