package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/IFT365/src/FinalPrj/customers"
	"github.com/IFT365/src/FinalPrj/dealers"
)

func main() {

	custFile, err := customers.ReadCustFile("customers.csv")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\n# of customers: %d - %s\n", len(custFile), custFile)

	fmt.Printf("%-4s\t%-20s\t%-20s\t%-15s\t%-s\t%-s\n",
		"ID",
		"Customer Name",
		"Street",
		"City",
		"State",
		"Zip",
	)

	for _, customer := range custFile {
		fmt.Printf("%s\t%-20s\t%-20s\t%-15s\t%-s\t%s\n",
			customer.CustomerId,
			customer.Name,
			customer.Address,
			customer.City,
			customer.State,
			customer.Zip,
		)
	}

	dealerFile, err := dealers.ReadDealerFile("dealers.csv")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\n# of dealers: %d - %s\n", len(dealerFile), dealerFile)

	serviceDate()

}

func serviceDate() {

	currentTime := time.Now()

	// Take the user input for a string
	fmt.Print("Enter a date to get a list of prior service: ")
	reader := bufio.NewReader(os.Stdin)
	inString, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Date entered: %s", inString)
	fmt.Println("Time entered: ", currentTime)
}
