package lang

type Employee struct {
	name   string
	salary int
	sales  int
	bonus  int
}

func NewEmployee(name string, salary, sales, bonus int) *Employee {
	var e Employee
	return newEmployee(&e, name, salary, sales, bonus)
}

func newEmployee(e *Employee, name string, salary, sales, bonus int) *Employee {
	e.name = name
	e.salary = salary
	e.sales = sales
	e.bonus = bonus
	return e
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

func Filter(employees []Employee, filter func(Employee) bool) []Employee {
	result := make([]Employee, 0)
	for _, e := range employees {
		if filter(e) {
			result = append(result, e)
		}
	}
	return result
}

// Return true if both slices contain the same employees
func Equals(xs, ys []Employee) bool {

	if len(xs) != len(ys) {
		return false
	}

	m := make(map[string]bool)

	for _, x := range xs {
		m[x.name] = true
	}

	for _, y := range ys {
		if _, exists := m[y.name]; !exists {
			return false
		}
	}

	return true

}

func FilterPointers(employees []*Employee, filter func(Employee) bool) []*Employee {
	result := make([]*Employee, 0)
	for _, pe := range employees {
		e := *pe
		if filter(e) {
			result = append(result, pe)
		}
	}
	return result
}

func EqualsPointers(xs, ys []*Employee) bool {

	if len(xs) != len(ys) {
		return false
	}

	m := make(map[string]bool)

	for _, x := range xs {
		m[(*x).name] = true
	}

	for _, y := range ys {
		if _, exists := m[(*y).name]; !exists {
			return false
		}
	}

	return true

}
