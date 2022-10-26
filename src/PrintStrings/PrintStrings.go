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
// entering strings the stored strings are printed out.
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// Settings that can be easily changed if needed.
const (
	ACCTNAMETEST string = "admin"    // The account to be matched
	ACCTPWTEST   string = "Pa$$w0rd" // The password to be matched

	LOGINATTEMPTMAX int = 3 // The number of allowed attempts
)

// Main procedure
func main() {

	moreStrings := true // Set the check to true to start
	myStrings := make([]string, 0)
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
		myStrings = append(myStrings, strings.TrimSpace(inString))

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
	fmt.Printf("%d strings entered\n", len(myStrings))

	fmt.Println(strings.Join(myStrings, ", "))

}
