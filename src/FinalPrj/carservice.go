package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/IFT365/src/FinalPrj/customers"
	"github.com/IFT365/src/FinalPrj/dealers"
	"github.com/IFT365/src/FinalPrj/technicians"
)

var CustomersList = make(map[string]customers.Customer)
var DealersList = make(map[string]dealers.Dealer)
var TechniciansList = make(map[string]technicians.Technician)

type CustData struct {
	CustomerCount int
	Customers     map[string]customers.Customer
	Dealers       map[string]dealers.Dealer
}

type CustViewData struct {
	Customer customers.Customer
	Dealer   dealers.Dealer
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func loadfiles() error {
	var err error
	CustomersList, err = customers.LoadCustomers("customers.csv", "")
	check(err)

	DealersList, err = dealers.LoadDealers("dealers.csv")
	check(err)

	TechniciansList, err = technicians.LoadTechnicians("technicians.csv")
	check(err)

	return err

}
func mainHandler(writer http.ResponseWriter, request *http.Request) {

	dealerID := request.URL.Query().Get("dealerID")

	CustomersList, err := customers.LoadCustomers("customers.csv", dealerID)
	check(err)

	html, err := template.ParseFiles("carservice.html")
	check(err)
	custData := CustData{
		CustomerCount: len(CustomersList),
		Customers:     CustomersList,
		Dealers:       DealersList,
	}
	err = html.Execute(writer, custData)
	check(err)
}

func custviewHandler(writer http.ResponseWriter, request *http.Request) {

	custID := request.URL.Query().Get("custID")

	err := loadfiles()
	check(err)

	html, err := template.ParseFiles("customerview.html")
	check(err)
	custviewdata := CustViewData{Customer: CustomersList[custID], Dealer: DealersList[CustomersList[custID].DealerID]}

	err = html.Execute(writer, custviewdata)
	check(err)
}

func serviceactionHandler(writer http.ResponseWriter, request *http.Request) {

	custID := request.URL.Query().Get("custID")

	if request.FormValue("OilChange") == "001" {
		fmt.Println("OC")
		row := fmt.Sprintf("%s,%s,%s,",
			time.Now().Format(time.RFC3339),
			custID,
			request.FormValue("OilChange"))

		options := os.O_WRONLY | os.O_APPEND | os.O_CREATE
		file, err := os.OpenFile("transactions.csv", options, os.FileMode(0600))
		check(err)
		_, err = fmt.Fprintln(file, row)
		check(err)
		err = file.Close()
		check(err)
	}

	if request.FormValue("CarWash") == "002" {
		fmt.Println("CW")
		row := fmt.Sprintf("%s,%s,%s",
			time.Now().Format(time.RFC3339),
			custID,
			request.FormValue("CarWash"))

		options := os.O_WRONLY | os.O_APPEND | os.O_CREATE
		file, err := os.OpenFile("transactions.csv", options, os.FileMode(0600))
		check(err)
		_, err = fmt.Fprintln(file, row)
		check(err)
		err = file.Close()
		check(err)
	}

}

func viewHandler(writer http.ResponseWriter, request *http.Request) {
	// signatures := getStrings("signatures.txt")
	html, err := template.ParseFiles("carservice.html")
	check(err)

	err = html.Execute(writer, nil)
	check(err)

}

func main() {

	err := loadfiles()
	check(err)

	http.HandleFunc("/carservice", mainHandler)
	//	http.HandleFunc("/customers", custHandler)
	http.HandleFunc("/customerview", custviewHandler)
	http.HandleFunc("/carservice/new", newcustomerHandler)
	http.HandleFunc("/carservice/addnew", newcustomerpostHandler)
	http.HandleFunc("/serviceaction", serviceactionHandler)

	err = http.ListenAndServe("localhost:8080", nil)
	log.Fatal(err)

}

// func newHandler displays a form to enter a signature.
func newcustomerHandler(writer http.ResponseWriter, request *http.Request) {
	html, err := template.ParseFiles("newcustomer.html")
	check(err)
	err = html.Execute(writer, nil)
	check(err)
}
func newcustomerpostHandler(writer http.ResponseWriter, request *http.Request) {

	fmt.Println("New Customer Writing")
	newID := len(CustomersList)

	row := fmt.Sprintf("%d,%s,%s,%s,%s,%s,%s,", newID,
		request.FormValue("Name"),
		request.FormValue("Address"),
		request.FormValue("City"),
		request.FormValue("State"),
		request.FormValue("Zip"),
		request.FormValue("dealer"))

	options := os.O_WRONLY | os.O_APPEND | os.O_CREATE
	file, err := os.OpenFile("customers.csv", options, os.FileMode(0600))
	check(err)
	_, err = fmt.Fprintln(file, row)
	check(err)
	err = file.Close()
	check(err)

	err = loadfiles()
	check(err)

	http.Redirect(writer, request, "/carservice", http.StatusFound)
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
