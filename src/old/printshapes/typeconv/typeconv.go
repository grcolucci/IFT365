package main

import (
	"fmt"
)

type Liters float64
type Gallons float64

func (l Liters) ToGallons() Gallons {
	return Gallons(l * 0.264)
}

func (g Gallons) ToLiters() Liters {
	return Liters(g * 3.785)
}

func main() {

	carFuel := Gallons(2)
	busFuel := Liters(50)

	fmt.Printf("%0.1f Gallons equals %0.1f liters\n", carFuel, carFuel.ToLiters())
	fmt.Printf("%0.1f Liters equals %0.1f gallons\n", busFuel, busFuel.ToGallons())

	// 	carFuel += ToGallons(Liters(8.0))
	// 	busFuel += ToLiters(Gallons(30.0))

	// 	fmt.Printf("%0.2f\n", carFuel)
	// 	fmt.Printf("%0.2f\n", busFuel)
	//
}
