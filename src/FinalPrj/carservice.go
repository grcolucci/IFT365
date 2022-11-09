package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
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

	// Enter the date
	serviceDate()

	// Print list of vehicles over 6 months since oil change

	// Print list of vehicles over 6 months since car wash

}

func serviceDate() {

	// Take the user input for a string
	fmt.Print("Enter the month (1-12): ")
	reader := bufio.NewReader(os.Stdin)
	inString, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	// Convert the salary to a float64
	monthIn, err := strconv.Atoi(strings.TrimSpace(inString))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print("Enter the day (1-31): ")
	reader = bufio.NewReader(os.Stdin)
	inString, err = reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	// Convert the salary to a float64
	dayIn, err := strconv.Atoi(strings.TrimSpace(inString))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print("Enter the year (1999-Present): ")
	reader = bufio.NewReader(os.Stdin)
	inString, err = reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	// Convert the salary to a float64
	yearIn, err := strconv.Atoi(strings.TrimSpace(inString))
	if err != nil {
		log.Fatal(err)
	}
	currentTime := time.Date(yearIn, time.Month(monthIn), dayIn, 0, 0, 0, 0, time.UTC)

	fmt.Println(yearIn, monthIn, dayIn)
	fmt.Println("Time entered: ", currentTime)
}
