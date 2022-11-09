package main

import "fmt"

type MyType string

func (m MyType) SayHi(s string) int {
	fmt.Println(s)
	fmt.Println(m)
	return 99
}

type Number int

func (n *Number) Double() {
	*n *= 2
}

func main() {

	value := MyType("testing")

	fmt.Println(value.SayHi("another"))

	number := Number(4)

	number.Double()

	fmt.Println(number)

}
