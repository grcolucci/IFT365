package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/IFT365/src/FinalPrj/customers"
)

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

	custID := request.URL.Query().Get("custID")
	customersList, err := customers.LoadCustomers("customers.csv")
	check(err)
	html, err := template.ParseFiles("customerview.html")
	check(err)

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
	http.HandleFunc("/customerview", custviewHandler)

	err := http.ListenAndServe("localhost:8080", nil)
	log.Fatal(err)

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
