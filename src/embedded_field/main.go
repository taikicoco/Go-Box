package main

import (
	"fmt"
)

type Employee struct { 
	Name string
	ID   string
}

func (e Employee) Description() string { 
	return fmt.Sprintf("%s (%s)", e.Name, e.ID)
}

type Manager struct {
	Employee // 型のみ書く（埋め込みフィールド） NameとIDが加わる
	Reports []Employee  // 部下（報告の対象者）  Employeeのスライス
}

func (m Manager) FindNewEmployees() []Employee {  
	newEmployees := []Employee{ // Employee（従業員）のスライス
		Employee{
			"石田三成",
			"13112",
		},
		Employee{
			"徳川家康",
			"13115",
		},
	}
	return newEmployees
}

type Inner struct { 
	X int
}

type Outer struct {
	Inner
	X int
}

func main() {
	m := Manager{
		Employee: Employee{
			Name: "豊臣秀吉",
			ID:   "12345",
		},
		Reports: []Employee{},
	}
	fmt.Println(m.ID) // 12345
	
	//同じ処理になる
	fmt.Println(m.Description()) // 豊臣秀吉 (12345)
	fmt.Println(m.Employee.Description())

	t := Employee{
		Name: "伊藤",
		ID: "982",
	}
	fmt.Println(t.Description())

	m.Reports = m.FindNewEmployees()
	fmt.Println(m.Employee)  // {豊臣秀吉 12345}
	fmt.Println(m.Reports)  // [{石田三成 13112} {徳川家康 13115}]

	o := Outer{ 
		Inner: Inner{
			X: 10,
		},
		X: 20,
	}
	fmt.Println(o.X)       // 20
	fmt.Println(o.Inner.X) // 10   
}
