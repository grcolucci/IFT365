package transactions

import (
	"encoding/csv"
	"fmt"
	"os"
	"sort"
	"strconv"
)

type Transaction struct {
	Date        string
	CustomerID  string
	ServiceType string
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

func LoadTransactions(fName string, sortBy SortList, filterBy FilterList) ([]Transaction, error) {

	csvFile, err := os.Open(fName)
	if err != nil {
		return nil, err
	}
	fmt.Println("Successfully Opened CSV file: ", fName)
	defer csvFile.Close()

	fmt.Println(sortBy)
	fmt.Println(filterBy)

	//	transactionList := make(map[string]Transaction)
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
			Technician:  line[3],
		}

		stat.Price, err = strconv.ParseFloat(line[4], 64)
		if err != nil {
			fmt.Println(err)
		}

		if filterBy.FilterBy == "custID" {
			if filterBy.FilterValue == stat.CustomerID {
				transactionsList = append(transactionsList, stat)
			}
		} else if filterBy.FilterBy == "custID" {
			if filterBy.FilterValue == stat.CustomerID {
				transactionsList = append(transactionsList, stat)
			}
		} else {
			transactionsList = append(transactionsList, stat)
		}

	}

	if sortBy.SortField == "date" {
		// Sort by last name
		sort.Slice(transactionsList, func(i, j int) bool {
			if sortBy.Ascending {
				return transactionsList[i].Date < transactionsList[j].Date
			} else {
				return transactionsList[i].Date > transactionsList[j].Date
			}
		})
	} else if sortBy.SortField == "type" {
		// Sort by last name
		sort.Slice(transactionsList, func(i, j int) bool {
			if sortBy.Ascending {
				return transactionsList[i].ServiceType < transactionsList[j].ServiceType
			} else {
				return transactionsList[i].ServiceType > transactionsList[j].ServiceType
			}
		})
	} else if sortBy.SortField == "tech" {
		// Sort by last name
		sort.Slice(transactionsList, func(i, j int) bool {
			if sortBy.Ascending {
				return transactionsList[i].Technician < transactionsList[j].Technician
			} else {
				return transactionsList[i].Technician > transactionsList[j].Technician
			}
		})
	} else if sortBy.SortField == "price" {
		// Sort by last name
		sort.Slice(transactionsList, func(i, j int) bool {
			if sortBy.Ascending {
				return transactionsList[i].Price < transactionsList[j].Price
			} else {
				return transactionsList[i].Price > transactionsList[j].Price
			}
		})
	}

	return transactionsList, nil

}
