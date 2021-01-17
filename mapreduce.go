package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

type Student struct {
	name string `json: "name"`
	age  int    `json: "age"`
}

func newStudent(name, age string) Student {
	a, err := strconv.Atoi(age)
	if err != nil {
		panic(err)
	}
	return Student{name, a}
}

func readLines(fileName string) [][]string {
	f, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		panic(err)
	}
	return lines
}

func reduceSumAge(students []Student) int {
	r := 0
	for _, s := range students {
		r += s.age
	}
	return r
}

// MapReduce show cases an example of map reduce pipeline
func MapReduce() {
	fileName := "students.csv"
	lines := readLines(fileName)
	students := make([]Student, 0, len(lines))
	fmt.Println(len(lines))
	for _, l := range lines {
		student := newStudent(l[0], l[1])
		students = append(students, student)
	}
	sumAge := reduceSumAge(students)
	fmt.Println(sumAge)
}
