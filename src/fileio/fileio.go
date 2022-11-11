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
package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"math/rand"
	"os"
	"sort"
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

// Teacher struct
type Teacher struct {
	EmployeeID int
	fName      string
	lName      string
	Salary     float64
	Address
}

// Method to Set and Get the teacher name
// Set passes in a pointer for the reciever so the name
// can truely be updated
func (t *Teacher) SetName(name string) {

	fields := strings.Split(name, " ")
	t.fName = fields[0]
	t.lName = fields[1]
}

func (t Teacher) GetName() string {

	return t.fName
}

// Function to print out address labels
// Input: Sudent/Teacher name field and Address struct/field
func PrintAddressLbl(t Teacher) {

	fmt.Printf("%-10s\t%-15s\t%-20s\t%s, %s %05d\n", t.fName, t.lName, t.Street, t.City, t.State, t.zipCode)

}

// Loop through the teacher list and add all the salaries together
func SalaryTotal(teachers []Teacher) float64 {

	var totSalary float64
	for _, v := range teachers {
		totSalary += v.Salary
	}

	return totSalary

}

func main() {

	// Open and read in the teacher file
	// The contents of the file are read into a slice of type Teacher.
	teachFile, err := ReadTeacherFile("src/fileio/teachers.csv")
	if err != nil {
		log.Fatal(err)
	}

	// Print out the contents of the slice and then formatted labels

	fmt.Println("Teacher slice contents")
	fmt.Println(teachFile)

	fmt.Printf("\n# of teachers loaded: %d\n", len(teachFile))

	fmt.Println("Teacher Labels:")
	for _, v := range teachFile {
		PrintAddressLbl(v)
	}
	sort.Slice(teachFile, func(i, j int) bool {
		return teachFile[i].lName < teachFile[j].lName
	})

	fmt.Println("Sorted Teacher Labels:")
	for _, v := range teachFile {
		PrintAddressLbl(v)
	}

	// Randomly select one of the teacher records and change the name.
	// Then print out a new label for that record.
	rand.Seed(time.Now().UnixNano())
	randIndex := rand.Intn((len(teachFile) - 1 + 1) + 0)
	teachFile[randIndex].SetName("Mickey Mouse")

	sort.Slice(teachFile, func(i, j int) bool {
		return teachFile[i].lName < teachFile[j].lName
	})

	fmt.Println("Updated and resorted teacher label")
	for _, v := range teachFile {
		PrintAddressLbl(v)
	}

	// Print out the total salary for all teachers
	fmt.Printf("The total salary for all teachers is: %0.2f", SalaryTotal(teachFile))

	// Write out the new data to a temp file

	file, err := os.Create("records.csv")

	if err != nil {
		log.Fatalln("failed to open file", err)
	}
	defer file.Close()
	w := csv.NewWriter(file)
	defer w.Flush()

	// Using WriteAll
	var data [][]string
	for _, record := range teachFile {
		row := []string{strconv.Itoa(record.EmployeeID), record.fName, record.lName}
		data = append(data, row)
	}
	w.WriteAll(data)

}

// Read in and store the teacher data
// Input: file name to be read
// Output: slice with all the teachers in it.
func ReadTeacherFile(fileName string) ([]Teacher, error) {

	// Open the file
	csvFile, err := os.Open(fileName)
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

		intVar, err := strconv.Atoi(line[0]) // Convert TeacherID string to int
		if err != nil {
			log.Panicf("Error with data record %d skipped", i)
			continue
		}
		stat.EmployeeID = intVar

		stat.fName = line[1]
		stat.lName = line[2]

		// Convert the salary to a float64
		fltVar, err := strconv.ParseFloat(line[3], 64)
		if err != nil {
			log.Panicf("Error with data record %d skipped", i)
			continue

		}
		stat.Salary = fltVar

		stat.Street = line[4]
		stat.City = line[5]
		stat.State = line[6]

		// Convert the zip code
		intVar, err = strconv.Atoi(line[7])
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
