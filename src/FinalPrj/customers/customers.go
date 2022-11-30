package customers

import (
	"encoding/csv"
	"fmt"
	"os"
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
	defer file.Close()

	r := csv.NewReader(file)

	records, err := r.ReadAll()
	if err != nil {
		return nil, err
	}

	customersList := make(map[string]Customer)

	for _, record := range records {

		if dealerID != "" {
			if record[7] != dealerID {
				continue
			}

		}

		cust := Customer{
			CustomerId: record[0],
			Name:       record[1],
			Address:    record[2],
			City:       record[3],
			State:      record[4],
			Zip:        record[5],
			Phone:      record[6],
			DealerID:   record[7],
		}
		cust.LastOilChange.ServiceDate = record[8]
		cust.LastOilChange.ServiceType = record[9]
		cust.LastOilChange.Dealer = record[10]
		cust.LastOilChange.Technician = record[11]

		cust.LastCarWash.ServiceDate = record[12]
		cust.LastCarWash.ServiceType = record[13]
		cust.LastCarWash.Dealer = record[14]
		cust.LastCarWash.Technician = record[15]

		cust.MenuLine = fmt.Sprintf("%10s%20s", cust.CustomerId, cust.Name)
		cust.URLLine = fmt.Sprintf("http://localhost:8080/carservice/customerview?custID=%s", cust.CustomerId)

		customersList[cust.CustomerId] = cust
	}

	return customersList, nil
}

func UpdateRecords(customersList map[string]Customer) error {

	fmt.Println("Update Customer File")

	var newRecs [][]string

	for _, cust := range customersList {
		rec := []string{
			cust.CustomerId,
			cust.Name,
			cust.Address,
			cust.City,
			cust.State,
			cust.Zip,
			cust.Phone,
			cust.DealerID,
			cust.LastOilChange.ServiceDate,
			cust.LastOilChange.ServiceType,
			cust.LastOilChange.Dealer,
			cust.LastOilChange.Technician,
			cust.LastCarWash.ServiceDate,
			cust.LastCarWash.ServiceType,
			cust.LastCarWash.Dealer,
			cust.LastCarWash.Technician,
		}

		newRecs = append(newRecs, rec)
	}

	options := os.O_WRONLY | os.O_CREATE
	file, err := os.OpenFile("customers.csv", options, os.FileMode(0600))
	if err != nil {
		return err
	}
	defer file.Close()

	w := csv.NewWriter(file)
	err = w.WriteAll(newRecs) // calls Flush internally
	if err != nil {
		return err
	}

	return nil

}

func AddRecord(newRec []string) error {

	fmt.Println(newRec)
	options := os.O_WRONLY | os.O_CREATE | os.O_APPEND
	file, err := os.OpenFile("customers.csv", options, os.FileMode(0600))
	if err != nil {
		return err
	}
	defer file.Close()

	w := csv.NewWriter(file)
	defer w.Flush()

	err = w.Write(newRec) // calls Flush internally
	if err != nil {
		return err
	}

	return nil

}
