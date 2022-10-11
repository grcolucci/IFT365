package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const ACCTNAMETEST string = "Glenn"
const ACCTPWTEST string = "pwtest"

const LOGINATTEMPTMAX int = 3

func main() {

	acctValid := false
	reader := bufio.NewReader(os.Stdin)

	for i := 0; i < LOGINATTEMPTMAX && !acctValid; i++ {
		fmt.Print("Enter Account: ")
		reader = bufio.NewReader(os.Stdin)
		acctName, err := reader.ReadString('\n')
		fmt.Println(acctName)
		fmt.Println(err)
		if err != nil {
			log.Fatal(err)
		}
		acctName = strings.TrimSpace(acctName)

		fmt.Print("Enter Password ")
		reader = bufio.NewReader(os.Stdin)
		acctPW, err := reader.ReadString('\n')
		acctPW = strings.TrimSpace(acctPW)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(acctPW)
		fmt.Println(err)

		if acctName == ACCTNAMETEST {
			fmt.Println("Account Match")
			if acctPW == ACCTPWTEST {
				acctValid = true
				fmt.Println("Match")
			}
		}
	}
}
