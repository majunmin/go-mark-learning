package interface_learn

import "fmt"

type SalaryCalculater interface {
	CalculateSalary() int
}

type Permanent struct {
	empId    int
	basicPay int
	pf       int
}

type Contract struct {
	empId    int
	basicPay int
}

// salary of Permanent employee has the basicPay and pf
func (p Permanent) CalculateSalary() int {
	return p.basicPay + p.pf
}

// salary of contract employee has the basicPay only
func (c Contract) CalculateSalary() int {
	return c.basicPay
}

/*
total expense is calculated by iterating through the salary slice and summing
the salaries of the individual employee
*/
func totalExpense(s []SalaryCalculater) {
	expense := 0
	for _, cal := range s {
		expense += cal.CalculateSalary()
	}
	fmt.Printf("Total Expense Per Monnth $%d", expense)
}
