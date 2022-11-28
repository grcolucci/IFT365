package services

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
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
