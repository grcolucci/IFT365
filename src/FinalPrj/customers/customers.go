// ////////////////////////////////////////////////////////////////////////////////////////
//
// Filename: customers.go
// File type: Go
// Author: Glenn Colucci
// Class: IFT 365
// Date: December 2, 2022
//
// Description: This file deals with the customer records and csv file
// All structures and operation pertaining to Customer records should be
// handled by this package.
package customers

import (
	"encoding/csv"
	"fmt"
	"os"
)

func (c Customer) ToSlice() []string {

	rec := []string{
		c.CustomerId,
		c.Name,
		c.Address,
		c.City,
		c.State,
		c.Zip,
		c.Phone,
		c.DealerID,
		c.Car1.Year,
		c.Car1.Brand,
		c.Car1.Model,
		c.Car1.LastCarWash.ServiceDate,
		c.Car1.LastCarWash.Dealer,
		c.Car1.LastCarWash.Technician,
		c.Car1.LastOilChange.ServiceDate,
		c.Car1.LastOilChange.Dealer,
		c.Car1.LastOilChange.Technician,
		c.Car2.Year,
		c.Car2.Brand,
		c.Car2.Model,
		c.Car2.LastCarWash.ServiceDate,
		c.Car2.LastCarWash.Dealer,
		c.Car2.LastCarWash.Technician,
		c.Car2.LastOilChange.ServiceDate,
		c.Car2.LastOilChange.Dealer,
		c.Car2.LastOilChange.Technician,
		c.LastOCPromo.PromoDate,
		c.LastCWPromo.PromoDate,
	}

	return rec

}

type Customer struct {
	CustomerId  string
	Name        string
	Address     string
	City        string
	State       string
	Zip         string
	Phone       string
	DealerID    string
	Car1        Car
	Car2        Car
	LastOCPromo LastPromo
	LastCWPromo LastPromo
	MenuLine    string
	URLLine     string
}

type LastTransaction struct {
	ServiceDate string
	Dealer      string
	Technician  string
}

type LastPromo struct {
	PromoDate string
}

type Car struct {
	Year          string
	Brand         string
	Model         string
	LastCarWash   LastTransaction
	LastOilChange LastTransaction
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

		cust.Car1.Year = record[8]
		cust.Car1.Brand = record[9]
		cust.Car1.Model = record[10]
		cust.Car1.LastCarWash.ServiceDate = record[11]
		cust.Car1.LastCarWash.Dealer = record[12]
		cust.Car1.LastCarWash.Technician = record[13]
		cust.Car1.LastOilChange.ServiceDate = record[14]
		cust.Car1.LastOilChange.Dealer = record[15]
		cust.Car1.LastOilChange.Technician = record[16]

		cust.Car2.Year = record[17]
		cust.Car2.Brand = record[18]
		cust.Car2.Model = record[19]
		cust.Car2.LastCarWash.ServiceDate = record[20]
		cust.Car2.LastCarWash.Dealer = record[21]
		cust.Car2.LastCarWash.Technician = record[22]
		cust.Car2.LastOilChange.ServiceDate = record[23]
		cust.Car2.LastOilChange.Dealer = record[24]
		cust.Car2.LastOilChange.Technician = record[25]
		cust.LastOCPromo.PromoDate = record[26]
		cust.LastCWPromo.PromoDate = record[27]

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
		rec := cust.ToSlice()
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
