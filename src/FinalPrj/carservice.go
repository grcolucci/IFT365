package main

import (
	"fmt"
	"log"

	"github.com/IFT365/src/FinalPrj/customers"
	"github.com/IFT365/src/FinalPrj/dealers"
)

func main() {

	custFile, err := customers.ReadCustFile("src/FinalPrj/customers.csv")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\n# of customers: %d - %s\n", len(custFile), custFile)

	fmt.Printf("%-4s\t%-20s\t%-25s\t%-15s\t%-s\t%-s\n",
		"ID",
		"Customer Name",
		"Street",
		"City",
		"State",
		"Zip",
	)

	for _, customer := range custFile {
		fmt.Printf("%s\t%-20s\t%-25s\t%-15s\t%-s\t\t%s\n",
			customer.CustomerId,
			customer.Name,
			customer.Address,
			customer.City,
			customer.State,
			customer.Zip,
		)
	}

	dealerFile, err := dealers.ReadDealerFile("src/FinalPrj/dealers.csv")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\n# of dealers: %d - %s\n", len(dealerFile), dealerFile)

}
