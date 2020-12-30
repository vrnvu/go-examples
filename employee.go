package main

type Employee struct {
	name   string
	salary int
	sales  int
	bonus  int
}

const BONUS_PERCENTAGE = 10

func getBonusPercentage(salary int) int {
	return (salary * BONUS_PERCENTAGE) / 100
}

func FindEmployeeBonus(salary, numberOfSales int) int {
	bonusPercentage := getBonusPercentage(salary)
	bonus := bonusPercentage * numberOfSales
	return bonus
}
