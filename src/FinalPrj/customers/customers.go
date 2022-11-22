package customers

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Customer struct {
	CustomerId    string
	Name          string
	Address       string
	City          string
	State         string
	Zip           string
	Phone         string
	MenuLine      string
	DealerID      string
	URLLine       string
	vehicles      []Car
	LastOilChange LastTransaction
	LastCarWash   LastTransaction
}

type LastTransaction struct {
	ServiceDate string
	ServiceType string
	Dealer      string
	Technician  string
}

type Car struct {
	CustomerId    int
	Name          string
	Year          int
	Model         string
	LastCarWash   string
	LastOilChange string
}

// LoadCustomers returns a slice of strings read from fileName, one
// string per line.
func LoadCustomers(fileName string, dealerID string) (map[string]Customer, error) {

	file, err := os.Open(fileName)
	if os.IsNotExist(err) {
		return nil, err
	}

	customersList := make(map[string]Customer)
	// check(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), ",")
		if dealerID != "" && dealerID != line[6] {
			continue
		}

		cust := Customer{
			CustomerId: line[0],
			Name:       line[1],
			Address:    line[2],
			City:       line[3],
			State:      line[4],
			Zip:        line[5],
			DealerID:   line[6],
		}
		cust.MenuLine = fmt.Sprintf("%10s%20s", cust.CustomerId, cust.Name)
		cust.URLLine = fmt.Sprintf("http://localhost:8080/customerview?custID=%s", cust.CustomerId)
		customersList[cust.CustomerId] = cust
	}
	// check()
	return customersList, scanner.Err()
}

func UpdateRecords(customerList map[string]Customer) error {

	fmt.Println("New Customer Writing")

	// options := os.O_WRONLY | os.O_APPEND | os.O_CREATE
	options := os.O_WRONLY | os.O_CREATE
	file, err := os.OpenFile("customers.csv", options, os.FileMode(0600))
	if err != nil {
		return err
	}

	for _, cust := range customerList {

		row := fmt.Sprintf("%s,%s,%s,%s,%s,%s,%s,%s,%s", cust.CustomerId,
			cust.Name,
			cust.Address,
			cust.City,
			cust.State,
			cust.Zip,
			cust.DealerID,
			cust.LastCarWash,
			cust.LastOilChange,
		)

		_, err = fmt.Fprintln(file, row)
		return err
	}

	err = file.Close()
	return err

}
