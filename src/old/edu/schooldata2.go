// ////////////////////////////////////////////////////////////////////////////////////////
//
// Filename: main.go
// File type: Go
// Author: Glenn Colucci
// Class: IFT 365
// Module: 4
// Date: October 31, 2022
//
// Description:
// Load information for teachers and students
// Load the info into a holding area and perform functions on the data
package main

import (
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

// Comman struct for teacher and student address
type Address struct {
	Street  string
	City    string
	State   string
	zipCode int
}

// Student struct
type Student struct {
	Studentid int
	Name      string
	Address
}

// Teacher struct
type Teacher struct {
	EmployeeID int
	Name       string
	Salary     float64
	Address
}

// Method to Set and Get the teacher name
// Set passes in a pointer for the reciever so the name
// can truely be updated
func (t *Teacher) SetName(name string) {

	t.Name = name
}

func (t Teacher) GetName() string {

	return t.Name
}

func main() {

	// Open and read in the teacher file
	// The contents of the file are read into a slice of type Teacher.
	teachFile, err := ReadTeacherFile()
	if err != nil {
		log.Fatal(err)
	}

	// Print out the contents of the slice and then formatted labels

	fmt.Println("Teacher slice contents")
	fmt.Println(teachFile)

	fmt.Printf("\n# of teachers loaded: %d\n", len(teachFile))

	fmt.Println("Teacher Labels:")
	for _, v := range teachFile {
		PrintAddressLbl(v.Name, v.Address)
	}

	// Open and read in the student file.
	// The contents of the file are read in and put into a list of type Student
	studentFile, err := ReadStudentFile()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\n# of students: %d\n\n", len(studentFile))

	// Print out the contents of the slice and then formated labels
	fmt.Println("Student slice contents")
	fmt.Println(studentFile)

	fmt.Println("Student Labels:")
	for _, v := range studentFile {
		PrintAddressLbl(v.Name, v.Address)
	}

	// Randomly select one of the teacher records and change the name.
	// Then print out a new label for that record.
	rand.Seed(time.Now().UnixNano())
	randIndex := rand.Intn((len(teachFile) - 1 + 1) + 0)
	teachFile[randIndex].SetName("Mickey Mouse")

	fmt.Println("Updated teacher label")

	PrintAddressLbl(teachFile[randIndex].Name, teachFile[randIndex].Address)

	// Print out the total salary for all teachers
	fmt.Printf("The total salary for all teachers is: %0.2f", SalaryTotal(teachFile))

}

// Function to print out address labels
// Input: Sudent/Teacher name field and Address struct/field
func PrintAddressLbl(inName string, inAddr Address) {

	fmt.Println(inName)
	fmt.Println(inAddr.Street)
	fmt.Printf("%s, %s %05d\n\n", inAddr.City, inAddr.State, inAddr.zipCode)
}

// Loop through the teacher list and add all the salaries together
func SalaryTotal(teachers []Teacher) float64 {

	var totSalary float64
	for _, v := range teachers {
		totSalary += v.Salary
	}

	return totSalary

}

// Read in and store the teacher data
// Input: file name to be read
// Output: slice with all the teachers in it.
func ReadTeacherFile() ([]Teacher, error) {

	teachLines := []string{"1,Glenn Colucci,49.97,43412 Countrywalk Ct.,Ashburn,VA,20147,",
		"2,Linda Colucci,89.34,34789 Alda Lane,Bethany Beach,DE,13390,",
		"3,Andrew Dinkins,37.87,394 Oak Lane,Leesburg,VA,2938,",
		"4,Sarah Putnam,32.98,394 Mountain Way,Rolling Hils,NC,2384,",
		"5,Teddy Bear,23.45,1 Main St.,Belmont,NC,39483,"}

	teacherFile := make([]Teacher, 0) // Create the teacher slice

	// Loop through each line read in and parse the data into
	// a slice element
	for i, line := range teachLines {

		var stat Teacher
		fields := strings.Split(line, ",")

		intVar, err := strconv.Atoi(fields[0]) // Convert TeacherID string to int
		if err != nil {
			log.Panicf("Error with data record %d skipped", i)
			continue
		}
		stat.EmployeeID = intVar

		stat.Name = fields[1]

		// Convert the salary to a float64
		fltVar, err := strconv.ParseFloat(fields[2], 64)
		if err != nil {
			log.Panicf("Error with data record %d skipped", i)
			continue

		}
		stat.Salary = fltVar

		stat.Street = fields[3]
		stat.City = fields[4]
		stat.State = fields[5]

		// Convert the zip code
		intVar, err = strconv.Atoi(fields[6])
		if err != nil {
			log.Panicf("Error with data record %d skipped", i)
			continue
		}

		stat.zipCode = intVar

		// Add new instance to the slice
		teacherFile = append(teacherFile, stat)

	}

	// Return the slice
	return teacherFile, nil

}

// Read in and store the student data
// Input: file name to be read
// Output: slice with all the students in it.

func ReadStudentFile() ([]Student, error) {

	studLines := [2]string{"1,Jennifer Colucci,22 Sunshine Ct.,St Petersburg,FL,33458,",
		"2,Julia Colucci,4839 Star Way,Hollywood,CA,04783,"}

	studentFile := make([]Student, 0)

	for i, line := range studLines {

		var stat Student

		fields := strings.Split(line, ",")

		intVar, err := strconv.Atoi(fields[0])
		if err != nil {
			log.Panicf("Error with data record %d skipped", i)
			continue
		}
		stat.Studentid = intVar

		stat.Name = fields[1]

		stat.Street = fields[2]
		stat.City = fields[3]
		stat.State = fields[4]
		intVar, err = strconv.Atoi(fields[5])
		if err != nil {
			log.Panicf("Error with data record %d skipped", i)
			continue
		}

		stat.zipCode = intVar

		studentFile = append(studentFile, stat)

	}

	return studentFile, nil

}
