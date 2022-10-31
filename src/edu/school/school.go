// ////////////////////////////////////////////////////////////////////////////////////////
//
// Filename: school.go
// File type: Go
// Author: Glenn Colucci
// Class: IFT 365
// Module: 4
// Date: October 31, 2022
//
// Description:
// Package for processing student and teacher data.
// storage structs and functions/methods for handling data
package school

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
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

// Mothod to Set and Get the teacher name
// Set passes in a pointer for the reciever so the name
// can truely be updated
func (t *Teacher) SetName(name string) {

	t.Name = name
}

func (t Teacher) GetName() string {

	return t.Name
}

// Function to print out address labels
// Input: Sudent/Teacher name field and Address struct/field
func PrintAddressLbl(inName string, inAddr Address) {

	fmt.Println(inName)
	fmt.Println(inAddr.Street)
	fmt.Printf("%s, %s %05d\n\n", inAddr.City, inAddr.State, inAddr.zipCode)
}

// Read in and store the teacher data
// Input: file name to be read
// Output: slice with all the teachers in it.
func ReadTeacherFile(fName string) ([]Teacher, error) {

	// Open the file
	csvFile, err := os.Open(fName)
	if err != nil {
		return nil, err
	}
	fmt.Println("Successfully Opened CSV file")
	defer csvFile.Close()

	teacherFile := make([]Teacher, 0) // Create the teacher slice

	csvLines, err := csv.NewReader(csvFile).ReadAll() // Read in all the records
	if err != nil {
		fmt.Println(err)
	}

	// Loop through each line read in and parse the data into
	// a slice element
	for i, line := range csvLines {

		var stat Teacher

		intVar, err := strconv.Atoi(line[0]) // Convert string to int
		if err != nil {
			log.Panicf("Error with data record %d skipped", i)
			continue
		}
		stat.EmployeeID = intVar

		stat.Name = line[1]

		// Convert the salary to a float64
		fltVar, err := strconv.ParseFloat(line[0], 64)
		if err != nil {
			log.Panicf("Error with data record %d skipped", i)
			continue

		}
		stat.Salary = fltVar

		stat.Street = line[3]
		stat.City = line[4]
		stat.State = line[5]

		// Convert the zip code
		intVar, err = strconv.Atoi(line[6])
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

func ReadStudentFile(fName string) ([]Student, error) {

	csvFile, err := os.Open(fName)
	if err != nil {
		return nil, err
	}
	fmt.Println("Successfully Opened CSV file")
	defer csvFile.Close()

	studentFile := make([]Student, 0)

	csvLines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		fmt.Println(err)
	}
	for i, line := range csvLines {

		var stat Student

		intVar, err := strconv.Atoi(line[0])
		if err != nil {
			log.Panicf("Error with data record %d skipped", i)
			continue
		}
		stat.Studentid = intVar

		stat.Name = line[1]

		stat.Street = line[2]
		stat.City = line[3]
		stat.State = line[4]
		intVar, err = strconv.Atoi(line[5])
		if err != nil {
			log.Panicf("Error with data record %d skipped", i)
			continue
		}

		stat.zipCode = intVar

		studentFile = append(studentFile, stat)

	}

	return studentFile, nil

}
