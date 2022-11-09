package main

import (
	"fmt"

	"github.com/IFT365/src/mypkg"
)

func main() {

	var value mypkg.MyInterface

	value = mypkg.MyType(5)

	value.MethodWithoutParameters()

	value.MethodWithParameters(123.5)

	fmt.Println(value.MethodWithReturnValue())

	fmt.Println(value)

}
