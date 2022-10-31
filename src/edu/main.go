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
// Read in 2 files of information for teachers and students
// Load the info into a holding area and perform functions on the data
package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/IFT365/src/edu/school"
)

func main() {

	// Open and read in the teacher file
	// The contents of the file are read into a slice of type Teacher.
	teachFile, err := school.ReadTeacherFile("src/edu/teachers.csv")
	if err != nil {
		log.Fatal(err)
	}

	// Print out the contents of the slice and then formatted labels

	fmt.Println("Teacher slice contents")
	fmt.Println(teachFile)

	fmt.Printf("\n# of teachers loaded: %d\n", len(teachFile))

	fmt.Println("Teacher Labels:")
	for _, v := range teachFile {
		school.PrintAddressLbl(v.Name, v.Address)
	}

	// Open and read in the student file.
	// The contents of the file are read in and put into a list of type Student
	studentFile, err := school.ReadStudentFile("src/edu/students.csv")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\n# of students: %d\n\n", len(studentFile))

	// Print out the contents of the slice and then formated labels
	fmt.Println("Student slice contents")
	fmt.Println(studentFile)

	fmt.Println("Student Labels:")
	for _, v := range studentFile {
		school.PrintAddressLbl(v.Name, v.Address)
	}

	// Randomly select one of the teacher records and change the name.
	// Then print out a new label for that record.
	rand.Seed(time.Now().UnixNano())
	randIndex := rand.Intn((len(teachFile) - 1 + 1) + 0)
	teachFile[randIndex].SetName("Mickey Mouse")

	fmt.Println("Updated teacher label")

	school.PrintAddressLbl(teachFile[randIndex].Name, teachFile[randIndex].Address)

}
