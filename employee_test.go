package main

import (
	"testing"
)

func assert(t *testing.T, got, want int) {
	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

type findEmployeeTest struct {
	arg1, arg2, arg3, expected int
}

type getBonusPercentageTest struct {
	arg1, expected int
}

var findEmployeeTests = []findEmployeeTest{
	findEmployeeTest{5000, 5, 0, 2500},
	findEmployeeTest{8500, 3, 0, 2550},
}

var getBonusPercentageTests = []getBonusPercentageTest{
	getBonusPercentageTest{100, 10},
	getBonusPercentageTest{864, 86},
}

func TestFindEmpoyeeBonus(t *testing.T) {
	for _, test := range findEmployeeTests {
		e := NewEmployee("Employee", test.arg1, test.arg2, test.arg3)
		got := FindEmployeeBonus(e.salary, e.sales)
		want := test.expected
		assert(t, got, want)
	}
}

func TestGetBonusPercentage(t *testing.T) {
	for _, test := range getBonusPercentageTests {
		got := getBonusPercentage(test.arg1)
		want := test.expected
		assert(t, got, want)
	}
}

func TestFilter(t *testing.T) {
	e0 := *NewEmployee("e0", 5000, 5, 0)
	e1 := *NewEmployee("e1", 6000, 5, 0)
	e2 := *NewEmployee("e2", 7000, 5, 0)
	employees := []Employee{e1, e0, e2}
	got := Filter(employees, func(e Employee) bool {
		return e.salary > 5500
	})
	want := []Employee{e2, e1}
	// want := []Employee{e2, e0}
	if !Equals(want, got) {
		t.Errorf("got %#v, wanted %#v", got, want)
	}
}
