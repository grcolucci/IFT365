// ////////////////////////////////////////////////////////////////////////////////////////
//
// Filename: PrintStrings.go
// File type: Go
// Author: Glenn Colucci
// Class: IFT 365
// Date: October 25, 2022
//
// Description:
// Users are asked to enters strings, one at a time and then prompted
// if they want to enter another.
// The entered strings are stored in a slice and once the user is finished
// entering strings, the stored strings are printed out.
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// Main procedure
func main() {

	moreStrings := true            // Set the check to true to start
	myStrings := make([]string, 0) // Slice to hold strings input

	// Loop through the user input
	//
	for moreStrings {

		// Take the user input for a string
		fmt.Print("Enter String: ")
		reader := bufio.NewReader(os.Stdin)
		inString, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		log.Println("String entered: ", inString)
		// Append string to end of slice.
		myStrings = append(myStrings, strings.TrimSpace(inString))

		// Check to see if user wants to input anothe string
		fmt.Print("Continue[Y/n]: ")
		reader = bufio.NewReader(os.Stdin)
		inString, err = reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		if strings.ToLower(strings.TrimSpace(inString)) != "y" {
			moreStrings = false
		}

	}

	// print out the number of strings entered
	fmt.Printf("%d strings entered\n", len(myStrings))

	// Print out all the strings entered.
	fmt.Println(strings.Join(myStrings, ", "))

}
