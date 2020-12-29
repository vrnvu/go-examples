package main

import (
	"errors"
	"math"
	"fmt"
)

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

// TODO compiler does TCO?
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

	// As pointers
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

// TODO in previous example is mentioned that a pass by value copies the struct
// How we would work with pointers only? A pointer to interface type?
type geometry interface {
	area() float64
}

type rectf64 struct {
	width, height float64
}

type circlef64 struct {
	radius float64
}

func (r *rectf64) area() float64 {
	return r.width * r.height
}

func (c circlef64) area() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry) {
	fmt.Println("&g", &g)
	fmt.Println(g)
	fmt.Println(g.area())
}

func Interfaces() {
	// TODO how I would pass pointer instead of values?
	r := rectf64{3, 4}
	c := circlef64{5}
	fmt.Println("&r", &r)
	fmt.Println("&c", &c)
	measure(&r)
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
	_, e := f2(42)
	if ae, ok := e.(*argError); ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.msg)
	}
}
