package school

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

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

func (t *Teacher) SetName(name string) {

	t.Name = name
}

func (t Teacher) GetName() string {

	return t.Name
}

func PrintAddressLbl(inName string, inAddr Address) {

	fmt.Println(inName)
	fmt.Println(inAddr.Street)
	fmt.Printf("%s, %s %5d\n\n", inAddr.City, inAddr.State, inAddr.zipCode)
}

func ReadTeacherFile(fName string) ([]Teacher, error) {

	csvFile, err := os.Open(fName)
	if err != nil {
		return nil, err
	}
	fmt.Println("Successfully Opened CSV file")
	defer csvFile.Close()

	teacherFile := make([]Teacher, 0)

	csvLines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		fmt.Println(err)
	}
	for i, line := range csvLines {

		var stat Teacher

		intVar, err := strconv.Atoi(line[0])
		if err != nil {
			log.Panicf("Error with data record %d skipped", i)
			continue
		}
		stat.EmployeeID = intVar

		stat.Name = line[1]
		fltVar, err := strconv.ParseFloat(line[0], 64)
		if err != nil {
			log.Panicf("Error with data record %d skipped", i)
			continue

		}
		stat.Salary = fltVar

		stat.Street = line[3]
		stat.City = line[4]
		stat.State = line[5]
		intVar, err = strconv.Atoi(line[6])
		if err != nil {
			log.Panicf("Error with data record %d skipped", i)
			continue
		}

		stat.zipCode = intVar

		teacherFile = append(teacherFile, stat)

	}

	fmt.Printf("\n# of teachers loaded: %d\n", len(teacherFile))

	for _, v := range teacherFile {
		PrintAddressLbl(v.Name, v.Address)
	}

	return teacherFile, nil

}

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

	fmt.Printf("\n# of Students loaded: %d\n", len(studentFile))

	for _, v := range studentFile {
		PrintAddressLbl(v.Name, v.Address)
	}

	return studentFile, nil

}
