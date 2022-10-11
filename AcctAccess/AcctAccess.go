package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	ACCTNAMETEST string = "Glenn"  // The account to be matched
	ACCTPWTEST   string = "pwtest" // The password to be matched

	LOGINATTEMPTMAX int = 3 // The number of allowed attempts
)

func main() {

	acctValid := false

	for i := 0; i < LOGINATTEMPTMAX && !acctValid; i++ {
		fmt.Print("Enter Account: ")
		reader := bufio.NewReader(os.Stdin)
		acctName, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		acctName = strings.TrimSpace(acctName)
		log.Println("Account entered: ", acctName)

		fmt.Print("Enter Password: ")
		reader = bufio.NewReader(os.Stdin)
		acctPW, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		acctPW = strings.TrimSpace(acctPW)
		log.Println("PW Entered: ", acctPW)

		if acctName == ACCTNAMETEST {
			log.Println("Account Match")
			if acctPW == ACCTPWTEST {
				acctValid = true
				log.Println("Password Match")
			} else {
				fmt.Println("The password entered is invalid.  Please try again.")
			}
		} else {
			fmt.Println("The account entered is invalid.  Please try again.")
		}
	}

	if acctValid {
		fmt.Println("\nYou have been logged in.")
	} else {
		fmt.Println("\nYou have exceeded the number of login attempts.  Exiting . . . ")
	}
}
