package dealers

import (
	"encoding/csv"
	"fmt"
	"os"
)

type Service struct {
	ID    string
	Name  string
	Price string
}

type Dealer struct {
	DealerID  string
	Name      string
	Telephone string
	Address   string
	City      string
	State     string
	Zip       string
	CarWash   Service
	OilChange Service
	MenuLine  string
	URLLine   string
}

func LoadDealers(fName string) (map[string]Dealer, error) {

	csvFile, err := os.Open(fName)
	if err != nil {
		return nil, err
	}
	fmt.Println("Successfully Opened CSV file:", fName)
	defer csvFile.Close()

	dealersList := make(map[string]Dealer)

	csvLines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		fmt.Println(err)
	}
	for _, line := range csvLines {
		stat := Dealer{
			DealerID:  line[0],
			Name:      line[1],
			Telephone: line[2],
			Address:   line[3],
			City:      line[4],
			State:     line[5],
			Zip:       line[6],
		}

		stat.OilChange.ID = "001"
		stat.OilChange.Name = "Premium Oil Change"
		stat.OilChange.Price = "$45.99"

		stat.CarWash.ID = "002"
		stat.CarWash.Name = "Full Service Car Wash"
		stat.CarWash.Price = "$10.99"

		stat.MenuLine = fmt.Sprintf("%10s%20s%15s", stat.DealerID, stat.Name, stat.Telephone)
		stat.URLLine = fmt.Sprintf("http://localhost:8080/carservice?dealerID=%s", stat.DealerID)

		dealersList[stat.DealerID] = stat

	}

	return dealersList, nil

}
