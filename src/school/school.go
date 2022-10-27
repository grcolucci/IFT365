package main

import "fmt"

type Address struct {
	Street  string
	City    string
	State   string
	zipCode int
}

type Student struct {
	Studentid int
	Name      string
	Address
}

type Teacher struct {
	EmployeeID int
	Name       string
	Salary     float64
	Address
}

func main() {

	teacherFile := make([]Teacher, 0)

	teach := Teacher{
		EmployeeID: 0001,
		Name:       "Mrs Whiggins",
		Salary:     73.99,
		// Adress {
		// 	Street: "test",
		// 	City: "city test",
		// 	State: "MM",
		// 	zipCode: 12345,
		// }
	}

	teach.Street = "245 Elm Street"
	teach.City = "Wayne"
	teach.State = "NJ"
	teach.zipCode = 12398

	teacherFile = append(teacherFile, teach)

	fmt.Println(teacherFile)

}
