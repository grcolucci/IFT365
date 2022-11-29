package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func main() {

	records := [][]string{
		{"first_name", "last_name", "occupation"},
		{"John", "Doe", "gardener"},
		{"Lucy", "Smith", "teacher"},
		{"Brian", "Bethamy", "programmer"},
		{"1", "2", "3", "4"},
		{"5", "6", "7", "8"},
	}

	f, err := os.Open("users.csv")

	if err != nil {

		log.Fatalln("failed to open file", err)
	}
	defer f.Close()

	w := csv.NewWriter(f)
	defer w.Flush()

	for _, record := range records {
		fmt.Println(record)
		if err := w.Write(record); err != nil {
			log.Fatalln("error writing record to file", err)
		}
	}

	rows := [][]string{
		{"1", "2", "3", "4"},
		{"5", "6", "7", "8"},
	}

	row1 := []string{"1", "2", "3", "4"}

	options := os.O_WRONLY | os.O_APPEND | os.O_CREATE
	file, err := os.OpenFile("transactions.csv", options, os.FileMode(0600))
	//	file, err := os.OpenFile("transactions.csv")
	defer f.Close()
	if err != nil {

		log.Fatalln("failed to open file", err)
	}

	u := csv.NewWriter(file)
	defer u.Flush()

	fmt.Println(rows[0])
	fmt.Println(rows[1])

	err = u.Write(row1)
	if err != nil {

		log.Fatalln("failed to open file", err)
	}

	// for _, row := range rows {
	// 	fmt.Println(row)
	// 	err = u.Write(row)
	// 	if err != nil {

	// 		log.Fatalln("failed to open file", err)
	// 	}
	// }

}
