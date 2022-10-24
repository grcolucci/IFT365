package stations

import (
	"encoding/csv"
	"fmt"
	"os"
)

type Service struct {
	serviceName  string
	servicePrice string
}

type Station struct {
	stationId string
	telephone string
	address   string
	city      string
	state     string
	zip       string
	carWash   Service
	oilChange Service
}

func ReadStationFile(fName string) ([]Station, error) {

	csvFile, err := os.Open(fName)
	if err != nil {
		return nil, err
	}
	fmt.Println("Successfully Opened CSV file")
	defer csvFile.Close()

	stationFile := make([]Station, 0)

	csvLines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		fmt.Println(err)
	}
	for _, line := range csvLines {
		stat := Station{
			stationId: line[0],
			telephone: line[1],
			address:   line[2],
			city:      line[3],
			state:     line[4],
			zip:       line[5],
		}

		stat.carWash.serviceName = "Full Service"
		stat.carWash.servicePrice = "$10.99"

		stat.oilChange.serviceName = "Premium"
		stat.oilChange.servicePrice = "$45.99"

		stationFile = append(stationFile, stat)

	}

	fmt.Println(stationFile)

	return stationFile, nil

}
