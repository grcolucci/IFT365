package main

import (
	"fmt"
	"log"

	"github.com/IFT365/src/gdates"
)

func main() {

	//date := gdates.Date{year: 2019, month: 5, day: 27}
	// fmt.Println(date)

	date2 := gdates.Date{}
	err := date2.SetYear(3)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("The years stored is: %d", date2.Year())
}
