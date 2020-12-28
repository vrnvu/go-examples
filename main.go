package main

import (
	"fmt"
	"strconv"
)

// Hello function returns a "Hello arg"
func Hello(arg int) string {
	result := "hello " + strconv.Itoa(arg)
	return result
}

// Returns a hard-coded value 11 for testing
func GetValueA() int {
	return 11
}

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(Hello(i))
	}
	// Similar to an If let expression
	if num := GetValueA(); num == 10 {
		fmt.Println("number was 10!")
	} else {
		fmt.Println("number was", num)
	}
}

//  LocalWords:  mv
