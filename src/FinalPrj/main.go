package main

import (
	"fmt"
	"log"

	"github.com/IFT365/src/FinalPrj/customers"
	"github.com/IFT365/src/FinalPrj/dealers"
)

func main() {

	custFile, err := customers.ReadCustFile("customers.csv")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\n# of customers: %d - %s\n", len(custFile), custFile)

	dealerFile, err := dealers.ReadDealerFile("dealers.csv")
	if err != nil {
		log.Fatal(err)
	}

}
