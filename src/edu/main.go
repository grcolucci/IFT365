package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/IFT365/src/edu/school"
)

func main() {

	teachFile, err := school.ReadTeacherFile("src/edu/teachers.csv")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\n# of teachers: %d\n", len(teachFile))

	studentFile, err := school.ReadStudentFile("src/edu/students.csv")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\n# of students: %d\n\n", len(studentFile))

	rand.Seed(time.Now().UnixNano())
	randIndex := rand.Intn((len(teachFile) - 1 + 1) + 0)
	teachFile[randIndex].SetName("Mickey Mouse")
	school.PrintAddressLbl(teachFile[randIndex].Name, teachFile[randIndex].Address)

}
