package transactions

import (
	"encoding/csv"
	"fmt"
	"os"
	"sort"
)

type Transaction struct {
	Date        string
	CustomerID  string
	ServiceType string
	Technician  string
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

		stat.MenuLine = fmt.Sprintf("%15s\t%10s\t%10s\t\t\t%10s",
			stat.Date,
			stat.CustomerID,
			stat.ServiceType,
			stat.Technician)

		//transactionsList[stat.ID] = stat
		if filterBy.FilterBy == "custID" {
			if filterBy.FilterValue == stat.CustomerID {
				transactionsList = append(transactionsList, stat)
			}
		} else {
			transactionsList = append(transactionsList, stat)
		}

	}

	if sortBy.SortField == "tech" {
		// Sort by last name
		sort.Slice(transactionsList, func(i, j int) bool {
			if sortBy.Ascending {
				return transactionsList[i].Technician < transactionsList[j].Technician
			} else {
				return transactionsList[i].Technician > transactionsList[j].Technician
			}
		})
	}

	return transactionsList, nil

}
