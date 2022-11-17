package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/IFT365/src/FinalPrj/customers"
)

/*
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
*/
func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
func custHandler(writer http.ResponseWriter, request *http.Request) {
	customersList, err := customers.LoadCustomers("customers.csv")
	check(err)
	html, err := template.ParseFiles("customers.html")
	check(err)
	custData := customers.CustData{
		CustomerCount: len(customersList),
		Customers:     customersList,
	}
	err = html.Execute(writer, custData)
	check(err)
}

func custviewHandler(writer http.ResponseWriter, request *http.Request) {

	custID := request.URL.Query().Get("cusID")
	customersList, err := customers.LoadCustomers("customers.csv")
	check(err)
	html, err := template.ParseFiles("customerView.html")
	check(err)
	//custData := customers.Customer

	//CustData{
	//	CustomerCount: len(customersList),
	//	Customers:     customersList,
	//}
	err = html.Execute(writer, customersList[custID])
	check(err)
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
	http.HandleFunc("/customerView", custviewHandler)

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
