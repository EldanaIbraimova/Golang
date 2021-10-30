package math

import {
	"math"
	"homework/salary"
}

func SumHourseOverworking (sal *salary.Salary) (float64) {
	oneHour := sal.BasicSalary / 100
	PayOverworking := sal.HourseOverworking * oneHour
	return PayOverworking
}

func SumTaxes (sal *salary.Salary) (float64) {
	salaryWithAllBonuses := AllBonuses(sal)
	pension := salaryWithAllBonuses / 10
	ipn := 0.00
	maxPension := 42500 * 50
	if pension >= maxPension {
		pension = maxPension
	}

	if salaryWithAllBonuses > 42500 {
		ipn = (salaryWithAllBonuses - pension - 42500) / 10
	}

	totalTaxes := pension + ipn
	return totalTaxes
}

func AllBonuses (sal *salary.Salary) (float64) {
	salaryWithAllBonuses := sal.BasicSalary + sal.Bonuses +  SumHourseOverworking(sal)
	return salaryWithAllBonuses
}

func totalSalary (sal *salary.Salary) {
salaryWithAllBonuses := AllBonuses(sal)
taxes := SumTaxes(sal)
totalSal := salaryWithAllBonuses - taxes
return totalSal
}