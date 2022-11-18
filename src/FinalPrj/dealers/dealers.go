package dealers

import (
	"encoding/csv"
	"fmt"
	"os"
)

type Service struct {
	serviceName  string
	servicePrice string
}

type Dealer struct {
	dealerID  string
	name      string
	telephone string
	address   string
	city      string
	state     string
	zip       string
	carWash   Service
	oilChange Service
	MenuLine  string
	URLLine   string
}

func LoadDealers(fName string) (map[string]Dealer, error) {

	csvFile, err := os.Open(fName)
	if err != nil {
		return nil, err
	}
	fmt.Println("Successfully Opened CSV file")
	defer csvFile.Close()

	dealers := make(map[string]Dealer)

	csvLines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		fmt.Println(err)
	}
	for _, line := range csvLines {
		stat := Dealer{
			dealerID:  line[0],
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
		stat.MenuLine = fmt.Sprintf("%10s%20s", stat.dealerID, stat.name)
		stat.URLLine = fmt.Sprintf("http://localhost:8080/customers?dealerID=%s", stat.dealerID)

		dealers[stat.dealerID] = stat

	}

	fmt.Printf("\n# of dealers loaded: %d\n", len(dealers))

	return dealers, nil

}
