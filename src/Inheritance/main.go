package main

import (
	"fmt"
)

type Employee struct { 
	Name string
	ID   string
}

type Manager struct {
	Employee
	Reports []Employee
} 

func (e Employee) Description() string {
	return fmt.Sprintf("%s (%s)", e.Name, e.ID)
}

func main() {
	m := Manager{ 
		Employee: Employee{
			Name: "上杉謙信",
			ID:   "12345",
		},
		Reports: []Employee{},
	}
	var eOK Employee = m.Employee // OK！
	fmt.Println(eOK)              // {上杉謙信 12345}
	//var eFail Employee = m        // コンパイル時のエラー！  

	o := Outer{
		Inner: Inner{
			A: 10,
		},
		S: "Hello",
	}
	fmt.Println(o.Double())  // Inner: 20
}

type Inner struct { 
	A int
}

func (i Inner) IntPrinter(val int) string {
	return fmt.Sprintf("Inner: %d", val)
}

func (i Inner) Double() string {
	result := i.A * 2
	return i.IntPrinter(result)
}

type Outer struct {
	Inner
	S string
}

func (o Outer) IntPrinter(val int) string {
	return fmt.Sprintf("Outer: %d", val)
}

