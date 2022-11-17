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
// load the data into a slice and write it out to a file
// read the data in from the just created file and sort
// the data in various ways, printing out each way
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

	fileName := "records.csv" // set the name of the file to be written/read

	// Load the data into a slice.
	teachFile, err := LoadTeacherFile()
	if err != nil {
		panic(err)
	}

	// Write out the data to a file
	err = WriteFile(teachFile, fileName)
	if err != nil {
		panic(err)
	}

	// read the data from the written file
	recordsFile, err := ReadTeacherFile(fileName)
	if err != nil {
		panic(err)
	}
	// Print out the contents of the slice

	fmt.Println("Teacher slice contents")
	fmt.Println(recordsFile)

	fmt.Printf("\n# of teachers loaded: %d\n", len(recordsFile))

	// Unsorted
	fmt.Println("\n\tTeacher List - Unsorted:")
	PrintAddressLbl(recordsFile)

	// Sort by last name
	sort.Slice(recordsFile, func(i, j int) bool {
		return recordsFile[i].lName < recordsFile[j].lName
	})

	fmt.Println("\n\tTeacher List - Sorted by last name:")
	PrintAddressLbl(recordsFile)

	// Randomly select one of the teacher records and change the name.
	// Then print out a new label for that record.
	rand.Seed(time.Now().UnixNano())
	randIndex := rand.Intn((len(recordsFile) - 1 + 1) + 0)
	recordsFile[randIndex].SetName("Mickey Mouse")

	// re-sort
	sort.Slice(recordsFile, func(i, j int) bool {
		return recordsFile[i].lName < recordsFile[j].lName
	})

	fmt.Println("\n\tTeacher list - One name changed sorted by last name")
	PrintAddressLbl(recordsFile)

	// Sort by Salary
	sort.Slice(recordsFile, func(i, j int) bool {
		return recordsFile[i].Salary > recordsFile[j].Salary
	})

	fmt.Println("\n\tTeacher list - sorted by salary (decending)")
	PrintAddressLbl(recordsFile)

}

func WriteFile(records []Teacher, fileName string) error {
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()
	w := csv.NewWriter(file)
	defer w.Flush()

	// Using WriteAll
	var data [][]string
	for _, record := range records {
		row := []string{strconv.Itoa(record.EmployeeID), record.fName, record.lName,
			fmt.Sprintf("%f", record.Salary), record.Street, record.City, record.State,
			strconv.Itoa(record.zipCode), ""}
		data = append(data, row)
	}
	w.WriteAll(data)

	fmt.Printf("Successfully created/wrote CSV file: %s\n", fileName)

	return nil

}

// load the teacher data
// Input: file name to be read
// Output: slice with all the teachers in it.
func LoadTeacherFile() ([]Teacher, error) {

	teacherFile := make([]Teacher, 0) // Create the teacher slice
	csvLines := []string{"1,Glenn,Colucci,49.97,43412 Countrywalk Ct.,Ashburn,VA,20147,",
		"2,Linda,Colucci,89.34,34789 Alda Lane,Bethany Beach,DE,13390,",
		"3,Andrew,Dinkins,37.87,394 Oak Lane,Leesburg,VA,2938,",
		"4,Sarah,Putnam,32.98,394 Mountain Way,Rolling Hils,NC,2384,",
		"5,Teddy,Bear,23.45,1 Main St.,Belmont,NC,39483,"}

	// Loop through each line read in and parse the data into
	// a slice element
	for i, line := range csvLines {

		splitline := strings.Split(line, ",")
		var stat Teacher

		intVar, err := strconv.Atoi(splitline[0]) // Convert TeacherID string to int
		if err != nil {
			log.Panicf("Error with data record %d skipped", i)
			continue
		}
		stat.EmployeeID = intVar

		stat.fName = splitline[1]
		stat.lName = splitline[2]

		// Convert the salary to a float64
		fltVar, err := strconv.ParseFloat(splitline[3], 64)
		if err != nil {
			log.Panicf("Error with data record %d skipped", i)
			continue

		}
		stat.Salary = fltVar

		stat.Street = splitline[4]
		stat.City = splitline[5]
		stat.State = splitline[6]

		// Convert the zip code
		intVar, err = strconv.Atoi(splitline[7])
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

// Read in and store the teacher data
// Input: file name to be read
// Output: slice with all the teachers in it.
func ReadTeacherFile(fileName string) ([]Teacher, error) {

	// Open the file
	csvFile, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	fmt.Printf("Successfully Opened CSV file: %s\n", fileName)
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
