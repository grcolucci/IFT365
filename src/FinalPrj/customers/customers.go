package customers

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

type Customer struct {
	CustomerId string
	Name       string
	Address    string
	City       string
	State      string
	Zip        string
	Phone      string
	MenuLine   string
	DealerID   string
	URLLine    string
	//	vehicles      []Car
	LastOilChange LastTransaction
	LastCarWash   LastTransaction
	LastOCPromo   LastPromo
	LastCWPromo   LastPromo
}

type LastTransaction struct {
	ServiceDate string
	ServiceType string
	Dealer      string
	Technician  string
}

type LastPromo struct {
	PromoDate time.Time
	PromoType string
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
		cust.LastOilChange.ServiceDate = line[7]
		cust.LastOilChange.ServiceType = line[8]
		cust.LastOilChange.Dealer = line[9]
		cust.LastOilChange.Technician = line[10]

		cust.LastCarWash.ServiceDate = line[11]
		cust.LastCarWash.ServiceType = line[12]
		cust.LastCarWash.Dealer = line[13]
		cust.LastCarWash.Technician = line[14]

		cust.MenuLine = fmt.Sprintf("%10s%20s", cust.CustomerId, cust.Name)
		cust.URLLine = fmt.Sprintf("http://localhost:8080/carservice/customerview?custID=%s", cust.CustomerId)
		customersList[cust.CustomerId] = cust
	}
	// check()
	return customersList, scanner.Err()
}

func UpdateRecords(customersList map[string]Customer) error {

	fmt.Println("Update Customer File")

	// options := os.O_WRONLY | os.O_APPEND | os.O_CREATE
	options := os.O_WRONLY | os.O_CREATE
	file, err := os.OpenFile("customers.csv", options, os.FileMode(0600))
	if err != nil {
		return err
	}

	for _, cust := range customersList {
		row := fmt.Sprintf("%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,", cust.CustomerId,
			cust.Name,
			cust.Address,
			cust.City,
			cust.State,
			cust.Zip,
			cust.DealerID,
			cust.LastOilChange.ServiceDate,
			cust.LastOilChange.ServiceType,
			cust.LastOilChange.Dealer,
			cust.LastOilChange.Technician,
			cust.LastCarWash.ServiceDate,
			cust.LastCarWash.ServiceType,
			cust.LastCarWash.Dealer,
			cust.LastCarWash.Technician,
		)

		_, err = fmt.Fprintln(file, row)
		if err != nil {
			return err
		}
	}

	err = file.Close()
	return err

}
