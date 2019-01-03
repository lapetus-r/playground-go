package main

import "fmt"

type Emp struct {
	Name   string
	Salary int
}

func main() {
	employees := make([]*Emp, 1) // Type앞에 *를 붙이면 포인터 선언
	employees[0] = &Emp{         // 값에 &를 붙이면 해당 값이 저장되는 주소를 나타냄
		Name:   "babokim",
		Salary: 1000,
	}
	for _, emp := range employees {
		emp.Salary = 2000
	}
	fmt.Println("After for each statement:", employees[0].Salary)
	for i := 0; i < len(employees); i++ {
		employees[i].Salary = 3000
	}
	fmt.Println("After for with index:", employees[0].Salary)
}
