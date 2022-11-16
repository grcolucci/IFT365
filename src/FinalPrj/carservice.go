package main

import (
	"bufio"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
)

type Customer struct {
	CustomerId string
	Name       string
	Address    string
	City       string
	State      string
	Zip        string
	Phone      string
	MenuLine   string
	//vehicles   []Car
}

type Car struct {
	CustomerId    int
	Name          string
	Year          int
	Model         string
	LastCarWash   string
	LastOilChange string
}
type CustData struct {
	CustomerCount int
	Customers     []Customer
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
func custHandler(writer http.ResponseWriter, request *http.Request) {
	customers := getStrings("customers.csv")
	html, err := template.ParseFiles("customers.html")
	check(err)
	custData := CustData{
		CustomerCount: len(customers),
		Customers:     customers,
	}
	err = html.Execute(writer, custData)
	check(err)
}

// getStrings returns a slice of strings read from fileName, one
// string per line.
func getStrings(fileName string) []Customer {

	var customers []Customer
	file, err := os.Open(fileName)
	if os.IsNotExist(err) {
		return nil
	}
	check(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), ",")
		cust := Customer{
			CustomerId: line[0],
			Name:       line[1],
			Address:    line[2],
			City:       line[3],
			State:      line[4],
			Zip:        line[5],
		}
		cust.MenuLine = fmt.Sprintf("%s\t%s\t", cust.CustomerId, cust.Name)
		customers = append(customers, cust)
	}
	check(scanner.Err())
	return customers
}

func viewHandler(writer http.ResponseWriter, request *http.Request) {
	// signatures := getStrings("signatures.txt")
	html, err := template.ParseFiles("carservice.html")
	check(err)
	//guestbook := Guestbook{
	//	SignatureCount: len(signatures),
	//	Signatures:     signatures,
	//}
	err = html.Execute(writer, nil)
	check(err)

}

func main() {

	http.HandleFunc("/carservice", viewHandler)
	http.HandleFunc("/customers", custHandler)

	err := http.ListenAndServe("localhost:8080", nil)
	log.Fatal(err)
	/* custFile, err := customers.ReadCustFile("customers.csv")
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
	*/
}

/*
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
*/
