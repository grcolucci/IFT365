// ////////////////////////////////////////////////////////////////////////////////////////
//
// Filename: services.go
// File type: Go
// Author: Glenn Colucci
// Class: IFT 365
// Date: December 2, 2022
//
// Description: This file deals with the services records and csv file
// All structures and operation pertaining to Service records should be
// handled by this package.
package services

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

// Settings that can be easily changed if needed.
const (
	OILCHANGE string = "001"
	CARWASH   string = "101"
)

type Service struct {
	ID    string
	Type  int
	Name  string
	Price float64
}

func LoadServices(fName string) (map[string]Service, error) {

	csvFile, err := os.Open(fName)
	if err != nil {
		return nil, err
	}
	fmt.Println("Successfully Opened CSV file:", fName)
	defer csvFile.Close()

	servicesList := make(map[string]Service)

	csvLines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		fmt.Println(err)
	}
	for _, line := range csvLines {
		stat := Service{
			ID:   line[0],
			Name: line[2],
		}
		stat.Type, err = strconv.Atoi(line[1])
		if err != nil {
			return nil, err
		}

		stat.Price, err = strconv.ParseFloat(line[3], 64)
		if err != nil {
			return nil, err
		}

		servicesList[stat.ID] = stat

	}

	return servicesList, nil

}
