package mypkg

import "fmt"

type MyInterface interface {
	MethodWithoutParameters()
	MethodWithParameters(float64)
	MethodWithReturnValue() string
}
type MyType int

func (m MyType) MethodWithoutParameters() {
	fmt.Println("MethodWithoutParameters")
}

func (m MyType) MethodWithParameters(f float64) {
	fmt.Println("MethodWithoutParameters with f ", f)
}

func (m MyType) MethodWithReturnValue() string {
	return "MethodWithReturnValue"
}

func (MyType) MethodNotInInterface() {
	fmt.Println("MethodNotInInterface")
}
