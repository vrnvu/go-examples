package main

import (
	"errors"
	"fmt"
	"math"
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

func ForIter() {
	for i := 0; i < 10; i++ {
		fmt.Println(Hello(i))
	}
}

func IfElseAndSwitch() {
	// Similar to an If let expression
	if num := GetValueA(); num == 10 {
		fmt.Println("number was 10!")
	} else {
		fmt.Println("number was", num)
	}

	// Same pattern than before but with a switch statement
	switch num := GetValueA(); num {
	case 10:
		fmt.Println("number was 10!")
	default:
		fmt.Println("number was", num)
	}
}

func Arrays() {
	var a [5]int
	fmt.Println("emp:", a)

	a[4] = 10
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])
	fmt.Println("len:", len(a))

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	var aa [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			aa[i][j] = i + j
		}
	}
	fmt.Println("2d:", aa)

	// Dynamic len of the array
	for i := 0; i < len(aa); i++ {
		for j := 0; j < len(aa[0]); j++ {
			aa[i][j] = i + j
		}
	}
	fmt.Println("2d:", aa)
}

func Slices() {
	// Slice is a container of type T
	// With size n
	// slice := make([]T, n)
	s := make([]string, 3)
	fmt.Println("emp:", s)

	// Set values
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("len:", len(s))
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	// We may get a new slice value from an append
	s = append(s, "d")
	fmt.Println("apd:", s)

	// Declare and initialize a
	// c slice with the same size and copy s
	c := make([]string, len(s))
	// What exactly copy does?
	// TODO
	copy(c, s)
	fmt.Println("cpy", c)

	// The copy does not move
	c[2] = "X"
	fmt.Println("org", s)
	fmt.Println("cpy", c)

	// Then we can slice
	l := s[1:4]
	fmt.Println("sl1", l)

	l = s[1:]
	fmt.Println("sl2", l)

	// Two dimensional slice
	// Notice we first initialize the external array
	ss := make([][]int, 3)
	for i := 0; i < len(ss); i++ {
		innerLen := i + 1
		// We declare the slice for ss[i]
		// notice the len of inner slices can vary
		ss[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			ss[i][j] = i + j
		}
	}
	fmt.Println("ss:", ss)

}

func Maps() {
	// Declare a map string => int
	m := make(map[string]int)
	// Initialize/Set some values
	m["k1"] = 1
	m["k2"] = 2
	fmt.Println("map:", m)

	// value, present
	// If value was present, present is true else false
	v1, prs1 := m["k1"]
	fmt.Println("v1:", v1)
	fmt.Println("prs1:", prs1)

	// To delete a key,value from a map
	delete(m, "k2")
	fmt.Println("map:", m)

	// If missing value has the zero value for int 0
	v2, prs2 := m["k2"]
	fmt.Println("v2:", v2)
	fmt.Println("prs2:", prs2)

	// Declare and initialize
	n := map[string]int{"k1": 1, "k2": 2}
	fmt.Println("n:", n)
}

func Ranges() {
	nums := []int{1, 2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	m := map[string]string{"k1": "1", "k2": "2"}
	for k, v := range m {
		fmt.Printf("%s -> %s\n", k, v)
	}
}

func Add(a, b, c int, s string) int {
	fmt.Println(s)
	return a + b + c
}

func AddMulti(a, b, c int, s string) (int, string) {
	return a + b + c, s
}

func Functions() {
	// Simple function
	result := Add(1, 2, 3, "Add called")
	fmt.Println(result)

	// Function with multiple return values
	r, s := AddMulti(1, 2, 3, "AddMulti called")
	fmt.Println("r: ", r, " s: ", s)

}

func VariadicFunctions(nums ...int) {
	fmt.Print(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func Closures() {
	nextInt := intSeq()

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
}

func RecursionFact(n int) int {
	if n == 0 {
		return 1
	}
	return n * RecursionFact(n-1)

}

func zeroval(ival int) {
	ival = 0
}

func zeroptr(iptr *int) {
	*iptr = 0
}

func Pointers() {
	i := 1
	fmt.Println("initial:", i)

	zeroval(i)
	fmt.Println("zeroval:", i)

	zeroptr(&i)
	fmt.Println("zeroptr:", i)

	fmt.Println("ptr:", &i)
}

type person struct {
	name string
	age  int
}

func newPerson(name string) *person {
	p := person{name: name}
	p.age = 42
	return &p
}

func Structs() {
	// Various ways to construct a person
	fmt.Println(person{"Bob", 20})
	fmt.Println(person{name: "Alice", age: 30})
	fmt.Println(person{name: "Fred"})

	// As references
	// & prefix yields a pointer to the struct
	fmt.Println(&person{name: "Ann", age: 40})

	// Its idiomatic to encapsulate new struct in constructor functions
	fmt.Println(newPerson("Jon"))

	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)
	fmt.Println(s.age)

	// Structs are mutable
	// If we copy by reference we will mutate both values
	// The pointers are automatically dereferenced
	sp := &s
	sp.age = 51
	fmt.Println(sp.age)
	fmt.Println(s.age)
}

type rect struct {
	width, height int
}

// receiver type *rect
func (r *rect) area() int {
	return r.width * r.height
}

// methods can be defined either pointer or value receiver types
func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

func Methods() {
	r := rect{10, 5}

	fmt.Println("area: ", r.area())
	fmt.Println("perim: ", r.perim())

	// Go automatically handles conversion between values and pointers for method calls
	// You may want to use a pointer receiver type to avoid copying on methods calls
	// Or to allow the method to mutate the receiving struct
	rp := &r
	fmt.Println("area: ", rp.area())
	fmt.Println("perim: ", rp.perim())
}

type geometry interface {
	area() float64
}

type rectf64 struct {
	width, height float64
}

type circlef64 struct {
	radius float64
}

func (r rectf64) area() float64 {
	return r.width * r.height
}

func (c circlef64) area() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
}

func Interfaces() {
	// TODO how I would pass references instead of values?
	r := rectf64{3, 4}
	c := circlef64{5}
	measure(r)
	measure(c)
}

// By conventions errors are the last value and have type error
func f1(arg int) (int, error) {
	if arg == 42 {
		return -1, errors.New("value was 42")
	}
	// A nil error indicates that there was no error
	return arg + 3, nil
}

// Custom Error by implementing Error interface
type argError struct {
	arg int
	msg string
}

func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.msg)
}

// Same function but using our custom error type
func f2(arg int) (int, error) {
	if arg == 42 {
		return -1, &argError{arg, "value was 42"}
	}
	return arg + 3, nil
}

func Errors() {
	for _, i := range []int{7, 42} {
		if r, e := f1(i); e != nil {
			fmt.Println("f1 failed:", e)
		} else {
			fmt.Println("f1 worked:", r)
		}
	}
	// Calling f2 instead, notice e print
	for _, i := range []int{7, 42} {
		if r, e := f2(i); e != nil {
			fmt.Println("f2 failed:", e)
		} else {
			fmt.Println("f2 worked:", r)
		}
	}

	// To use the data inside our custom error struct
	_, e:= f2(42)
	if ae, ok := e.(*argError); ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.msg)
	}
}

func main() {
	// ForIter()
	// IfElseAndSwitch()
	// Arrays()
	// Slices()
	// Maps()
	// Ranges()
	// Functions()
	// VariadicFunctions(1, 2)
	// VariadicFunctions(1, 2, 3)
	// VariadicFunctions([]int{1, 2, 3, 4}...)
	// Closures()
	// fmt.Println(RecursionFact(7))
	// Pointers()
	// Structs()
	// Methods()
	// Interfaces()
	// Errors()
}

//  LocalWords:  mv
