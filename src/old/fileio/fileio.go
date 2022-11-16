// ////////////////////////////////////////////////////////////////////////////////////////
//
// Filename: fileio.go
// File type: Go
// Author: Glenn Colucci
// Class: IFT 365
// Module: 5
// Date: November 15, 2022
//
// Description:
// Read in a list of data from a file "teachers.csv"
// sort the data in various ways, printing out each way
// Write the data out to a file "records.csv"
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

// struct for address
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

// Catch the panic
func reportPanic() {
	p := recover()

	if p == nil {
		return
	}

	err, ok := p.(error)
	if ok {
		fmt.Println(err)
	}
}

// Method to Set and Get the teacher name
// Set passes in a pointer for the reciever so the name
// can truely be updated
func (t *Teacher) SetName(name string) {

	fields := strings.Split(name, " ")
	t.fName = fields[0]
	t.lName = fields[1]
}

// Function to print out teacher list
func PrintAddressLbl(teacherList []Teacher) {

	fmt.Printf("\n%-10s\t%-15s\t%s\t%-20s\t%s, %s %s\n", "First Name", "Last Name",
		"Salary", "Street", "City", "St", "Zip Code")

	for _, t := range teacherList {

		fmt.Printf("%-10s\t%-15s\t%0.2f\t%-20s\t%s, %s %05d\n", t.fName, t.lName,
			t.Salary, t.Street, t.City, t.State, t.zipCode)
	}
}

func main() {

	defer reportPanic()

	// Open and read in the teacher file
	// The contents of the file are read into a slice of type Teacher.
	teachFile, err := ReadTeacherFile("teachers.csv")
	if err != nil {
		panic(err)
	}

	// Print out the contents of the slice

	fmt.Println("Teacher slice contents")
	fmt.Println(teachFile)

	fmt.Printf("\n# of teachers loaded: %d\n", len(teachFile))

	// Unsorted
	fmt.Println("\n\tTeacher List - Unsorted:")
	PrintAddressLbl(teachFile)

	// Sort by last name
	sort.Slice(teachFile, func(i, j int) bool {
		return teachFile[i].lName < teachFile[j].lName
	})

	fmt.Println("\n\tTeacher List - Sorted by last name:")
	PrintAddressLbl(teachFile)

	// Randomly select one of the teacher records and change the name.
	// Then print out a new label for that record.
	rand.Seed(time.Now().UnixNano())
	randIndex := rand.Intn((len(teachFile) - 1 + 1) + 0)
	teachFile[randIndex].SetName("Mickey Mouse")

	// re-sort
	sort.Slice(teachFile, func(i, j int) bool {
		return teachFile[i].lName < teachFile[j].lName
	})

	fmt.Println("\n\tTeacher list - One name changed sorted by last name")
	PrintAddressLbl(teachFile)

	// Sort by Salary
	sort.Slice(teachFile, func(i, j int) bool {
		return teachFile[i].Salary > teachFile[j].Salary
	})

	fmt.Println("\n\tTeacher list - sorted by salary (decending)")
	PrintAddressLbl(teachFile)

	// Write out the new data to a  file
	file, err := os.Create("records.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	w := csv.NewWriter(file)
	defer w.Flush()

	// Using WriteAll
	var data [][]string
	for _, record := range teachFile {
		row := []string{strconv.Itoa(record.EmployeeID), record.fName, record.lName,
			fmt.Sprintf("%f", record.Salary), record.Street, record.City, record.State,
			strconv.Itoa(record.zipCode), ""}
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
