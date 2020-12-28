package main

import ("fmt")

// To execute run
// $ go run *.go
func main() {
	ForIter()
	IfElseAndSwitch()
	Arrays()
	Slices()
	Maps()
	Ranges()
	Functions()
	VariadicFunctions(1, 2)
	VariadicFunctions(1, 2, 3)
	VariadicFunctions([]int{1, 2, 3, 4}...)
	Closures()
	fmt.Println(RecursionFact(7))
	Pointers()
	Structs()
	Methods()
	Interfaces()
	Errors()
}

//  LocalWords:  mv
