package technicians

import (
	"encoding/csv"
	"fmt"
	"os"
)

type Technician struct {
	ID   string
	Name string
}

func LoadTechnicians(fName string) (map[string]Technician, error) {

	csvFile, err := os.Open(fName)
	if err != nil {
		return nil, err
	}
	fmt.Println("Successfully Opened CSV file: ", fName)
	defer csvFile.Close()

	techniciansList := make(map[string]Technician)

	csvLines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		fmt.Println(err)
	}
	for _, line := range csvLines {
		stat := Technician{
			ID:   line[0],
			Name: line[1],
		}

		techniciansList[stat.ID] = stat

	}

	return techniciansList, nil

}
