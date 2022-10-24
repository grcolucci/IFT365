package customers

import (
	"encoding/csv"
	"fmt"
	"os"
)

type Customer struct {
	customerId string
	firstName  string
	lastName   string
	address    string
	city       string
	state      string
	zip        string
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
			customerId: line[0],
			firstName:  line[1],
			lastName:   line[2],
			address:    line[3],
			city:       line[4],
			state:      line[5],
			zip:        line[6],
		}

		custFile = append(custFile, cust)

	}

	fmt.Println(custFile)

	return custFile, nil

}
