package main

import "fmt"

// type Employee struct {
// 	name     string
// 	age      int
// 	division string
// }

type Person struct {
	name string
	age  int
}

// type Employee struct {
// 	person   Person
// 	division string
// }

func main() {
	// var employee Employee
	// employee.name = "raditeo"
	// employee.age = 23
	// employee.division = "test"

	// fmt.Println(employee.name)
	// fmt.Println(employee.age)
	// fmt.Println(employee.division)

	// var employee1 = Employee{}

	// employee1.name = "Raditeo"
	// employee1.age = 23
	// employee1.division = "test"

	// var employee2 = Employee{name: "Warma", age: 24, division: "test"}

	// fmt.Printf("employee1 %+v\n", employee1)
	// fmt.Printf("employee2 %+v\n", employee2)

	// var employee1 = Employee{name: "Warma", age: 24, division: "test"}

	// var employee2 *Employee = &employee1

	// fmt.Println("employee1 name:", employee1.name)
	// fmt.Println("employee2 name:", employee2.name)

	// fmt.Println(strings.Repeat("#", 25))

	// employee2.name = "Raditeo"

	// fmt.Println("employee1 name:", employee1.name)
	// fmt.Println("employee2 name:", employee2.name)

	// var employee1 = Employee{}
	// employee1.person.name = "Raditeo"
	// employee1.person.age = 23
	// employee1.division = "Test"

	// fmt.Printf("%+v", employee1)

	// var employee1 = struct {
	// 	person   Person
	// 	division string
	// }{}

	// employee1.person = Person{name: "Raditeo", age: 23}
	// employee1.division = "Test"

	// var employee2 = struct {
	// 	person   Person
	// 	division string
	// }{
	// 	person:   Person{name: "Warma", age: 23},
	// 	division: "Test",
	// }

	// fmt.Printf("employee1: %+v\n", employee1)
	// fmt.Printf("employee2: %+v\n", employee2)

	var people = []Person{
		{name: "raditeo", age: 23},
		{name: "warma", age: 23},
		{name: "loheh", age: 23},
	}

	for _, v := range people {
		fmt.Printf("%+v\n", v)
	}
}
