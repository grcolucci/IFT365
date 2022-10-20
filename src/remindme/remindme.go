package main

import (
	"fmt"
	"log"

	"github.com/IFT365/src/gdates"
)

func main() {

	date := gdates.Date{Year: 2019, Month: 5, Day: 27}
	fmt.Println(date)

	date2 := gdates.Date{}
	err := date2.SetYear(-3)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(date2.Year)
}
