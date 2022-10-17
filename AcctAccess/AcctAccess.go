// ////////////////////////////////////////////////////////////////////////////////////////
//
// Filename: AcctAccess.go
// File type: Go
// Author: Glenn Colucci
// Class: IFT 365
// Date: October 18, 2022
//
// Description: AcctAccess is an app that checks a users account and password.  Users can
// attempt a certain number of tries to gain access.
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

// Function to check the account and password input by the user
// Inputs: Account(String), Password(String)
// Returns: Account Valid(bool), Msg(reason if validation failed - nil if succesful)
func acctValidCheck(uName string, uPW string) (bool, error) {

	if uName == ACCTNAMETEST { // Test the username input
		log.Println("Account Match")
		if uPW == ACCTPWTEST { // Test the password input
			log.Println("Password Match")
			return true, nil // Account and PW are valid, no error msg
		} else {
			log.Println("Invlaid password")
			return false, fmt.Errorf("password incorrect")
		}
	} else {
		log.Println("Invalid Account")
		return false, fmt.Errorf("account incorrect")
	}
}

// Main procedure
func main() {

	acctValid := false
	err := fmt.Errorf("")

	// Loop through the user input
	//
	for i := 1; i <= LOGINATTEMPTMAX && !acctValid; i++ {

		// Take the user input for Account
		fmt.Print("Enter Account: ")
		reader := bufio.NewReader(os.Stdin)
		acctName, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		acctName = strings.ToLower(strings.TrimSpace(acctName)) // remove caps and the \n
		log.Println("Account entered: ", acctName)

		// Take the user input for the Password
		fmt.Print("Enter Password: ")
		reader = bufio.NewReader(os.Stdin)
		acctPW, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		acctPW = strings.TrimSpace(acctPW) // Remove the \n
		log.Println("PW Entered: ", acctPW)

		// Call the procedure to check the account and pw input
		acctValid, err = acctValidCheck(acctName, acctPW)
		if !acctValid {
			fmt.Println(acctValid)
			fmt.Println(err)
			if i < LOGINATTEMPTMAX { // Are there any attempts left?
				fmt.Printf(" - Please try again (%d Attempts left).\n", LOGINATTEMPTMAX-i)

			}
		}

	}
	fmt.Println(err)
	// Output the final outcome
	if acctValid {
		fmt.Println("\nLogin Successful!")
	} else {
		fmt.Println("\nYou have exceeded the number of login attempts.  Account locked – Contact 800-123-4567")
	}
}
