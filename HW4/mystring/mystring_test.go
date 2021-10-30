package mystring

import(
	"homework/salary"
	"testing"
)

func TestCheckWorker (t *testing.T){
	worker1 := salary.Salary{
		Name: "Aliya",
	    BasicSalary: 200000.00,
	    HourseOverworking : 10,
	    Bonuses float64 : 50000.00
} 
	}
}

isExist := CheckWorker(worker1.Name)
if isExist == false {
	t.Fail()
}