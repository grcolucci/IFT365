package transactions

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

type Transaction struct {
	Date        string
	CustomerID  string
	ServiceType string
	CarNum      string
	Technician  string
	Price       float64
	MenuLine    string
}

type SortList struct {
	SortField string
	Ascending bool
}

type FilterList struct {
	FilterBy    string
	FilterValue string
}

func LoadTransactions(fName string) ([]Transaction, error) {

	csvFile, err := os.Open(fName)
	if err != nil {
		return nil, err
	}
	fmt.Println("Successfully Opened CSV file: ", fName)
	defer csvFile.Close()

	transactionsList := make([]Transaction, 0)

	csvLines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		fmt.Println(err)
	}
	for _, line := range csvLines {
		stat := Transaction{
			Date:        line[0],
			CustomerID:  line[1],
			ServiceType: line[2],
			CarNum:      line[3],
			Technician:  line[4],
		}

		stat.Price, err = strconv.ParseFloat(line[5], 64)
		if err != nil {
			fmt.Println(err)
		}

		transactionsList = append(transactionsList, stat)
	}

	return transactionsList, nil

}

func WriteTransactions(rows [][]string) error {

	options := os.O_WRONLY | os.O_APPEND | os.O_CREATE
	file, err := os.OpenFile("transactions.csv", options, os.FileMode(0600))

	if err != nil {
		return err
	}

	defer file.Close()

	w := csv.NewWriter(file)
	defer w.Flush()

	err = w.WriteAll(rows)
	if err != nil {
		return err
	}

	return nil
}
