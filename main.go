package main

import (
	"fmt"
	"strconv"
)

// Hello function returns a "Hello arg"
func Hello(arg int) string {
	return "hello " + strconv.Itoa(arg)
}

func main() {
	for i := 0; i < 10; i++{
		fmt.Println(Hello(i))
	}
}

//  LocalWords:  mv
