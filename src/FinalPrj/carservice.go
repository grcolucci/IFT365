package main

import (
	"fmt"
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"sort"
	"strconv"
	"time"

	"github.com/IFT365/src/FinalPrj/customers"
	"github.com/IFT365/src/FinalPrj/dealers"
	"github.com/IFT365/src/FinalPrj/promotions"
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
	ServiceDate  string
	CustomerName string
	DealerName   string
	ServiceType  string
	Car          string
	TechName     string
	Price        string
}

type PromoMgmt struct {
	DaysPrior     string
	CustomerCnt   int
	ListStatus    string
	Customers     []customers.Customer
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
	sortBy.Ascending, _ = strconv.ParseBool(request.FormValue("order"))

	filterBy.FilterBy = request.FormValue("filterField")
	filterBy.FilterValue = request.FormValue("filterValue")

	TransactionsList, err := transactions.LoadTransactions("transactions.csv")
	check(err)

	var tdList []TransactionsDisplayList
	for _, tran := range TransactionsList {

		dLine := TransactionsDisplayList{
			ServiceDate:  tran.Date,
			CustomerName: CustomersList[tran.CustomerID].Name,
			ServiceType:  ServicesList[tran.ServiceType].Name,
			TechName:     TechniciansList[tran.Technician].Name,
		}

		if tran.CarNum == "1" {
			dLine.Car = fmt.Sprintf("%s %s %s",
				CustomersList[tran.CustomerID].Car1.Year,
				CustomersList[tran.CustomerID].Car1.Brand,
				CustomersList[tran.CustomerID].Car1.Model,
			)
		} else {
			dLine.Car = fmt.Sprintf("%s %s %s",
				CustomersList[tran.CustomerID].Car2.Year,
				CustomersList[tran.CustomerID].Car2.Brand,
				CustomersList[tran.CustomerID].Car2.Model,
			)
		}
		dLine.Price = fmt.Sprintf("%0.2f", tran.Price)

		if filterBy.FilterBy == "cust" {
			if filterBy.FilterValue == dLine.CustomerName {
				tdList = append(tdList, dLine)
			}
		} else if filterBy.FilterBy == "date" {
			if filterBy.FilterValue == dLine.ServiceDate {
				tdList = append(tdList, dLine)
			}
		} else if filterBy.FilterBy == "type" {
			if filterBy.FilterValue == dLine.ServiceType {
				tdList = append(tdList, dLine)
			}
		} else if filterBy.FilterBy == "dealer" {
			if filterBy.FilterValue == dLine.DealerName {
				tdList = append(tdList, dLine)
			}
		} else {
			tdList = append(tdList, dLine)
		}
	}

	if sortBy.SortField == "date" {
		// Sort by last name
		sort.Slice(tdList, func(i, j int) bool {
			if sortBy.Ascending {
				return tdList[i].ServiceDate < tdList[j].ServiceDate
			} else {
				return tdList[i].ServiceDate > tdList[j].ServiceDate
			}
		})
	} else if sortBy.SortField == "type" {
		// Sort by last name
		sort.Slice(tdList, func(i, j int) bool {
			if sortBy.Ascending {
				return tdList[i].ServiceType < tdList[j].ServiceType
			} else {
				return tdList[i].ServiceType > tdList[j].ServiceType
			}
		})
	} else if sortBy.SortField == "tech" {
		// Sort by last name
		sort.Slice(tdList, func(i, j int) bool {
			if sortBy.Ascending {
				return tdList[i].TechName < tdList[j].TechName
			} else {
				return tdList[i].TechName > tdList[j].TechName
			}
		})
	} else if sortBy.SortField == "price" {
		// Sort by last name
		sort.Slice(tdList, func(i, j int) bool {
			if sortBy.Ascending {
				return tdList[i].Price < tdList[j].Price
			} else {
				return tdList[i].Price > tdList[j].Price
			}
		})
	}

	html, err := template.ParseFiles("transactions.html")
	check(err)

	err = html.Execute(writer, tdList)
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
		CustomerId:  custID,
		Name:        CustomersList[custID].Name,
		Address:     CustomersList[custID].Address,
		City:        CustomersList[custID].City,
		State:       CustomersList[custID].State,
		Zip:         CustomersList[custID].Zip,
		Phone:       CustomersList[custID].Phone,
		DealerID:    CustomersList[custID].DealerID,
		Car1:        CustomersList[custID].Car1,
		Car2:        CustomersList[custID].Car2,
		LastOCPromo: CustomersList[custID].LastOCPromo,
		LastCWPromo: CustomersList[custID].LastCWPromo,
	}

	var rows [][]string

	if request.FormValue("service001") == "001" {

		updateCustRec = true
		if request.FormValue("servicecar1") == "1" {
			fmt.Println("Car1 Oil")
			row := []string{time.Now().Format("01-02-2006"),
				custID,
				services.OILCHANGE,
				"1",
				TechniciansList[fmt.Sprintf("%03d", randIndex)].ID,
				fmt.Sprintf("%f", ServicesList[services.OILCHANGE].Price),
			}

			rows = append(rows, row)

			cust.Car1.LastOilChange.Dealer = CustomersList[custID].DealerID
			cust.Car1.LastOilChange.ServiceDate = time.Now().Format("01-02-2006")
			cust.Car1.LastOilChange.Technician = TechniciansList[fmt.Sprintf("%03d", randIndex)].ID

		}

		if request.FormValue("servicecar2") == "2" {
			fmt.Println("Car2 Oil")
			row := []string{time.Now().Format("01-02-2006"),
				custID,
				services.OILCHANGE,
				"2",
				TechniciansList[fmt.Sprintf("%03d", randIndex)].ID,
				fmt.Sprintf("%f", ServicesList[services.OILCHANGE].Price),
			}

			rows = append(rows, row)

			cust.Car2.LastOilChange.Dealer = CustomersList[custID].DealerID
			cust.Car2.LastOilChange.ServiceDate = time.Now().Format("01-02-2006")
			cust.Car2.LastOilChange.Technician = TechniciansList[fmt.Sprintf("%03d", randIndex)].ID

		}
	}

	if request.FormValue("service101") == services.CARWASH {
		updateCustRec = true
		if request.FormValue("servicecar1") == "1" {
			fmt.Println("Car1 Wash")
			row := []string{time.Now().Format("01-02-2006"),
				custID,
				"101",
				"1",
				TechniciansList[fmt.Sprintf("%03d", randIndex)].ID,
				fmt.Sprintf("%f", ServicesList[services.CARWASH].Price),
			}

			rows = append(rows, row)

			cust.Car1.LastCarWash.Dealer = CustomersList[custID].DealerID
			cust.Car1.LastCarWash.ServiceDate = time.Now().Format("01-02-2006")
			cust.Car1.LastCarWash.Technician = TechniciansList[fmt.Sprintf("%03d", randIndex)].ID

		}

		if request.FormValue("servicecar2") == "2" {
			fmt.Println("Car2 Wash")
			row := []string{time.Now().Format("01-02-2006"),
				custID,
				services.CARWASH,
				"2",
				TechniciansList[fmt.Sprintf("%03d", randIndex)].ID,
				fmt.Sprintf("%f", ServicesList[services.CARWASH].Price),
			}

			rows = append(rows, row)

			cust.Car2.LastCarWash.Dealer = CustomersList[custID].DealerID
			cust.Car2.LastCarWash.ServiceDate = time.Now().Format("01-02-2006")
			cust.Car2.LastCarWash.Technician = TechniciansList[fmt.Sprintf("%03d", randIndex)].ID

		}

	}

	if updateCustRec {

		err := transactions.WriteTransactions(rows)
		check(err)

		// Delete customer recod from list, then re-add updated record

		delete(CustomersList, custID)

		CustomersList[cust.CustomerId] = cust

		// Write out the updated list to the file
		err = customers.UpdateRecords(CustomersList)
		check(err)
	}

	http.Redirect(writer, request, fmt.Sprintf("http://localhost:8080/carservice/customerview?custID=%s", custID), http.StatusFound)

}

/* func viewHandler(writer http.ResponseWriter, request *http.Request) {

	html, err := template.ParseFiles("carservice.html")
	check(err)

	err = html.Execute(writer, nil)
	check(err)

} */

func promomgmtHandler(writer http.ResponseWriter, request *http.Request) {

	html, err := template.ParseFiles("promomgmt.html")
	check(err)

	promomgmt := PromoMgmt{
		DisableButton: true,
	}
	var promoList []customers.Customer

	promomgmt.DaysPrior = request.FormValue("lastpromodays")
	promoType := request.FormValue("promotype")

	if len(promomgmt.DaysPrior) > 0 {
		promoList, err = promotions.PromoSelections(
			promomgmt.DaysPrior, promoType, CustomersList)
	}
	promomgmt.CustomerCnt = len(promoList)
	promomgmt.Customers = promoList
	promomgmt.ListStatus = fmt.Sprintf("Seletect Records for over %s days past and for %s type.", promomgmt.DaysPrior, ServicesList[promoType].Name)
	if promomgmt.CustomerCnt > 0 {
		promomgmt.DisableButton = false
	}

	err = html.Execute(writer, promomgmt)
	check(err)

}

func promosendHandler(writer http.ResponseWriter, request *http.Request) {

	promoType := request.URL.Query().Get("promotype")
	promoDays := request.URL.Query().Get("promoDays")

	if len(promoDays) > 0 {
		promoList, err := promotions.PromoSelections(
			promoDays, promoType, CustomersList)
		if err != nil {
			check(err)
		}

		for _, pCust := range promoList {
			custID := pCust.CustomerId
			newCust := customers.Customer{
				CustomerId: CustomersList[custID].CustomerId,
				Name:       CustomersList[custID].Name,
				Address:    CustomersList[custID].Address,
				City:       CustomersList[custID].City,
				State:      CustomersList[custID].State,
				Zip:        CustomersList[custID].Zip,
				Phone:      CustomersList[custID].Phone,
				DealerID:   CustomersList[custID].DealerID,
			}

			newCust.Car1.Year = CustomersList[custID].Car1.Year
			newCust.Car1.Brand = CustomersList[custID].Car1.Brand
			newCust.Car1.Model = CustomersList[custID].Car1.Model

			newCust.Car1.LastCarWash = CustomersList[custID].Car1.LastCarWash
			newCust.Car1.LastOilChange = CustomersList[custID].Car1.LastOilChange

			newCust.Car2.Year = CustomersList[custID].Car2.Year
			newCust.Car2.Brand = CustomersList[custID].Car2.Brand
			newCust.Car2.Model = CustomersList[custID].Car2.Model
			newCust.Car2.LastCarWash = CustomersList[custID].Car2.LastCarWash
			newCust.Car2.LastOilChange = CustomersList[custID].Car2.LastOilChange
			newCust.LastOCPromo = CustomersList[custID].LastOCPromo
			newCust.LastCWPromo = CustomersList[custID].LastCWPromo

			if promoType == "1" {
				newCust.LastOCPromo.PromoDate = time.Now().Format("2006-01-02")

			} else {
				newCust.LastCWPromo.PromoDate = time.Now().Format("2006-01-02")
			}
			delete(CustomersList, custID)

			CustomersList[custID] = newCust

		}
	}

	err := customers.UpdateRecords(CustomersList)
	check(err)

	err = loadfiles()
	check(err)

	http.Redirect(writer, request, fmt.Sprintf("http://localhost:8080/carservice"), http.StatusFound)

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
		DealerID:   request.FormValue("Dealer"),
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
