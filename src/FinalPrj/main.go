package main

import (
	"fmt"
	"log"

	"github.com/IFT365/src/FinalPrj/customers"
	"github.com/IFT365/src/FinalPrj/stations"
)

func main() {

	custFile, err := customers.ReadCustFile("customers.csv")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\n# of customers: %d - %s\n", len(custFile), custFile)

	statFile, err := stations.ReadStationFile("stations.csv")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\n# of stations: %d - %s\n", len(statFile), statFile)
}
