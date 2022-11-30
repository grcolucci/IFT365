package main

import (
	"encoding/csv"
	"fmt"
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/IFT365/src/FinalPrj/customers"
	"github.com/IFT365/src/FinalPrj/dealers"
	"github.com/IFT365/src/FinalPrj/services"
	"github.com/IFT365/src/FinalPrj/technicians"
	"github.com/IFT365/src/FinalPrj/transactions"
)

var CustomersList = make(map[string]customers.Customer)
var DealersList = make(map[string]dealers.Dealer)
var TechniciansList = make(map[string]technicians.Technician)
var ServicesList = make(map[string]services.Service)
var TransactionsList = make([]transactions.Transaction, 0)

type CustData struct {
	CustomerCount int
	Customers     map[string]customers.Customer
	Dealers       map[string]dealers.Dealer
}

type CustViewData struct {
	Customer    customers.Customer
	Dealer      dealers.Dealer
	ServicesCnt int
	Services    map[string]services.Service
	Technician  technicians.Technician
}

type TransactionsDisplayList struct {
	serviceDate string
	DealerName  string
	ServiceType string
	TechName    string
}

type PromoMgmt struct {
	DaysPrior     string
	CustomerCnt   int
	Customers     map[string]customers.Customer
	DisableButton bool
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

	ServicesList, err = services.LoadServices("services.csv")
	check(err)
	//TransactionsList, err = transactions.LoadTransactions("transactions.csv")
	//check(err)

	return err

}

func transactionListHandler(writer http.ResponseWriter, request *http.Request) {

	var sortBy transactions.SortList
	var filterBy transactions.FilterList

	sortBy.SortField = request.FormValue("sortby")
	sortBy.Ascending, _ = strconv.ParseBool(request.URL.Query().Get("order"))

	filterBy.FilterBy = request.URL.Query().Get("filterField")
	filterBy.FilterValue = request.URL.Query().Get("filterValue")

	TransactionsList, err := transactions.LoadTransactions("transactions.csv", sortBy, filterBy)
	check(err)
	var tm string
	for i, tran := range TransactionsList {
		tm = fmt.Sprintf("%15s\t%10s\t%10s\t\t\t%10s\t%0.2f",
			tran.Date,
			tran.CustomerID,
			ServicesList[tran.ServiceType].Name,
			TechniciansList[tran.Technician].Name,
			tran.Price)
		nl := &TransactionsList[i]
		(*nl).MenuLine = tm

	}

	html, err := template.ParseFiles("transactions.html")
	check(err)

	err = html.Execute(writer, TransactionsList)
	check(err)
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
	custviewdata := CustViewData{Customer: CustomersList[custID],
		Dealer:      DealersList[CustomersList[custID].DealerID],
		ServicesCnt: len(ServicesList),
		Services:    ServicesList}

	err = html.Execute(writer, custviewdata)
	check(err)
}

func serviceactionHandler(writer http.ResponseWriter, request *http.Request) {

	updateCustRec := false
	custID := request.URL.Query().Get("custID")

	rand.Seed(time.Now().UnixNano())
	randIndex := rand.Intn((len(TechniciansList) - 1 + 1) + 0)
	if randIndex <= 0 {
		randIndex = 1
	}

	cust := customers.Customer{
		CustomerId: custID,
		Name:       CustomersList[custID].Name,
		Address:    CustomersList[custID].Address,
		City:       CustomersList[custID].City,
		State:      CustomersList[custID].State,
		Zip:        CustomersList[custID].Zip,
		Phone:      CustomersList[custID].Phone,
		DealerID:   CustomersList[custID].DealerID,
		Car1:       CustomersList[custID].Car1,
		Car2:       CustomersList[custID].Car2,
		//LastOilChange: CustomersList[custID].LastOilChange,
		//LastCarWash:   CustomersList[custID].LastCarWash,
	}

	if request.FormValue("service001") == "001" {

		updateCustRec = true

		row := []string{time.Now().Format("01-02-2006"),
			custID,
			"001",
			TechniciansList[fmt.Sprintf("%03d", randIndex)].ID,
			fmt.Sprintf("%f", ServicesList["001"].Price),
		}

		// cust.LastOilChange.Dealer = CustomersList[custID].LastOilChange.Dealer
		// cust.LastOilChange.ServiceDate = time.Now().Format("01-02-2006")
		// cust.LastOilChange.ServiceType = "001"
		// fmt.Printf("Rand %03d\n", randIndex)
		// cust.LastOilChange.Technician = TechniciansList[fmt.Sprintf("%03d", randIndex)].ID

		options := os.O_WRONLY | os.O_APPEND | os.O_CREATE
		file, err := os.OpenFile("transactions.csv", options, os.FileMode(0600))

		if err != nil {
			log.Fatalln("failed to open file", err)
		}

		defer file.Close()

		w := csv.NewWriter(file)
		defer w.Flush()

		err = w.Write(row)
		check(err)
	}

	if request.FormValue("service101") == "101" {

		updateCustRec = true

		fmt.Println("CW")

		row := []string{time.Now().Format("01-02-2006"),
			custID,
			"101",
			TechniciansList[fmt.Sprintf("%03d", randIndex)].ID,
			fmt.Sprintf("%f", ServicesList["101"].Price),
		}

		options := os.O_WRONLY | os.O_APPEND | os.O_CREATE
		file, err := os.OpenFile("transactions.csv", options, os.FileMode(0600))

		if err != nil {

			log.Fatalln("failed to open file", err)
		}

		defer file.Close()

		w := csv.NewWriter(file)
		defer w.Flush()

		err = w.Write(row)
		check(err)

		// cust.LastCarWash.Dealer = CustomersList[custID].LastCarWash.Dealer
		// cust.LastCarWash.ServiceDate = time.Now().Format("01-02-2006")
		// cust.LastCarWash.ServiceType = request.FormValue("101")
		// cust.LastCarWash.Technician = TechniciansList[fmt.Sprintf("%03d", randIndex)].ID
	}

	if updateCustRec {
		// Delete customer recod from list, then re-add updated record

		delete(CustomersList, custID)

		CustomersList[cust.CustomerId] = cust

		// Write out the updated list to the file
		err := customers.UpdateRecords(CustomersList)
		check(err)
	}

	http.Redirect(writer, request, fmt.Sprintf("http://localhost:8080/carservice/customerview?custID=%s", custID), http.StatusFound)

}

func viewHandler(writer http.ResponseWriter, request *http.Request) {

	html, err := template.ParseFiles("carservice.html")
	check(err)

	err = html.Execute(writer, nil)
	check(err)

}

func promomgmtHandler(writer http.ResponseWriter, request *http.Request) {

	html, err := template.ParseFiles("promomgmt.html")
	check(err)

	var promomgmt PromoMgmt

	promomgmt.CustomerCnt = len(CustomersList)
	promomgmt.Customers = CustomersList
	if promomgmt.CustomerCnt > 1 {
		promomgmt.DisableButton = false
	}

	err = html.Execute(writer, promomgmt)
	check(err)

}

func promosendHandler(writer http.ResponseWriter, request *http.Request) {

	custID := request.URL.Query().Get("custID")
	http.Redirect(writer, request, fmt.Sprintf("http://localhost:8080/carservice/customerview?custID=%s", custID), http.StatusFound)

}

func main() {

	err := loadfiles()
	check(err)

	http.HandleFunc("/carservice", mainHandler)
	//	http.HandleFunc("/customers", custHandler)
	http.HandleFunc("/carservice/customerview", custviewHandler)
	http.HandleFunc("/carservice/updatecustomer", updatecustHandler)

	http.HandleFunc("/carservice/new", newcustomerHandler)
	http.HandleFunc("/carservice/addnew", newcustomerpostHandler)
	http.HandleFunc("/carservice/serviceaction", serviceactionHandler)
	http.HandleFunc("/carservice/transactionsview", transactionListHandler)

	http.HandleFunc("/carservice/promomgmt", promomgmtHandler)
	http.HandleFunc("/carservice/promosend", promosendHandler)

	err = http.ListenAndServe("localhost:8080", nil)
	log.Fatal(err)

}

// func newHandler displays a form to enter a signature.
func newcustomerHandler(writer http.ResponseWriter, request *http.Request) {
	html, err := template.ParseFiles("newcustomer.html")
	check(err)
	err = html.Execute(writer, DealersList)
	check(err)
}
func newcustomerpostHandler(writer http.ResponseWriter, request *http.Request) {

	fmt.Println("New Customer Writing")
	newID := len(CustomersList)

	cust := customers.Customer{
		CustomerId: fmt.Sprintf("%d", newID),
		Name:       request.FormValue("Name"),
		Address:    request.FormValue("Address"),
		City:       request.FormValue("City"),
		State:      request.FormValue("State"),
		Zip:        request.FormValue("Zip"),
		Phone:      request.FormValue("Phone"),
		DealerID:   request.FormValue("dealer"),
	}

	cust.Car1.Year = request.FormValue("car1year")
	cust.Car1.Brand = request.FormValue("car1brand")
	cust.Car1.Model = request.FormValue("car1model")
	cust.Car2.Year = request.FormValue("car2year")
	cust.Car2.Brand = request.FormValue("car2brand")
	cust.Car2.Model = request.FormValue("car2model")

	err := customers.AddRecord((cust.ToSlice()))
	check(err)

	err = loadfiles()
	check(err)

	http.Redirect(writer, request, fmt.Sprintf("http://localhost:8080/carservice/customerview?custID=%d", newID), http.StatusFound)
}

func updatecustHandler(writer http.ResponseWriter, request *http.Request) {

	fmt.Println("Update Customer Writing")
	custID := request.FormValue("CustomerID")

	newCust := customers.Customer{
		CustomerId: custID,
		Name:       request.FormValue("Name"),
		Address:    request.FormValue("Address"),
		City:       request.FormValue("City"),
		State:      request.FormValue("State"),
		Zip:        request.FormValue("Zip"),
		Phone:      request.FormValue("Phone"),
		DealerID:   request.FormValue("dealer"),
	}

	newCust.Car1.Year = request.FormValue("car1year")
	newCust.Car1.Brand = request.FormValue("car1brand")
	newCust.Car1.Model = request.FormValue("car1model")

	newCust.Car1.LastCarWash = CustomersList[custID].Car1.LastCarWash
	newCust.Car1.LastOilChange = CustomersList[custID].Car1.LastOilChange

	newCust.Car2.Year = request.FormValue("car2year")
	newCust.Car2.Brand = request.FormValue("car2brand")
	newCust.Car2.Model = request.FormValue("car2model")
	newCust.Car2.LastCarWash = CustomersList[custID].Car2.LastCarWash
	newCust.Car2.LastOilChange = CustomersList[custID].Car2.LastOilChange
	newCust.LastOCPromo = CustomersList[custID].LastOCPromo
	newCust.LastCWPromo = CustomersList[custID].LastCWPromo

	delete(CustomersList, custID)

	CustomersList[custID] = newCust

	err := customers.UpdateRecords(CustomersList)
	check(err)

	err = loadfiles()
	check(err)

	http.Redirect(writer, request, fmt.Sprintf("http://localhost:8080/carservice/customerview?custID=%s", custID), http.StatusFound)
}
