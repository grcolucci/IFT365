// ////////////////////////////////////////////////////////////////////////////////////////
//
// Filename: printStrings.go
// File type: Go
// Author: Glenn Colucci
// Class: IFT 365
// Date: October 18, 2022
//
// Description:
// attempt a certain number of tries to gain access.
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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

	// Loop through the user input
	//
	for moreStrings {

		// Take the user input for Account
		fmt.Print("Enter Account: ")
		reader := bufio.NewReader(os.Stdin)
		inString, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		
		log.Println("String entered: ", inString)	
		
		fmt.Print("Continue[Y/n]: ")
		reader := bufio.NewReader(os.Stdin)
		inString, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		if strings..tolower().TrimSpace(inString) != 'y' {
			moreStrings = false
		}

	}
}