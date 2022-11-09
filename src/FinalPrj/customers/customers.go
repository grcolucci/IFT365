package customers

import (
	"encoding/csv"
	"fmt"
	"os"
)

type Customer struct {
	CustomerId string
	Name       string
	Address    string
	City       string
	State      string
	Zip        string
	Phone      string
	vehicles   []Car
}

type Car struct {
	CustomerId    int
	Name          string
	Year          int
	Model         string
	LastCarWash   string
	LastOilChange string
}

func ReadCustFile(fName string) ([]Customer, error) {

	csvFile, err := os.Open(fName)
	if err != nil {
		return nil, err
	}
	fmt.Println("Successfully Opened CSV file")
	defer csvFile.Close()

	custFile := make([]Customer, 0)

	csvLines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		fmt.Println(err)
	}
	for _, line := range csvLines {
		cust := Customer{
			CustomerId: line[0],
			Name:       line[1],
			Address:    line[2],
			City:       line[3],
			State:      line[4],
			Zip:        line[5],
		}

		custFile = append(custFile, cust)

	}

	fmt.Println(custFile)

	return custFile, nil

}
