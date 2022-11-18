package customers

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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

// LoadCustomers returns a slice of strings read from fileName, one
// string per line.
func LoadCustomers(fileName string, dealerID string) (map[string]Customer, error) {

	customers := make(map[string]Customer)
	file, err := os.Open(fileName)
	if os.IsNotExist(err) {
		return customers, err
	}
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
		customers[cust.CustomerId] = cust
	}
	// check()
	return customers, scanner.Err()
}
