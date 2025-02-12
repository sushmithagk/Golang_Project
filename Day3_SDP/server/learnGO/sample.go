package main

import (
	"fmt"
)

type Car struct {
	Id     int
	Number string
	Mode1  string
	Type   string
}

func main() {
	//car1 :=Car {Id:101,Number:"KA4642",Mode1:"toyata",Type:"SUV"}
	var car1 Car = Car{Id: 101, Number: "KA4642", Mode1: "toyata", Type: "SUV"}
	fmt.Println(car1)
	fmt.Println(car1.Number)

	var cars []Car = []Car{
		{Id: 101, Number: "KA4642", Mode1: "toyata", Type: "SUV"},
		{Id: 25, Number: "TN4642", Mode1: "honda", Type: "SUV"},
	}
	fmt.Println(cars)

	//     fmt.Println("Hello, World!")
	// 	 first :=10
	// 	 sec :=20
	// 	 sum :=first+sec
	// 	fmt.Println(sum)
	// 	//var salaries[] float32=[] float32{1.0,2.0}
	// 	salaries :=[] float32{1.0,2.222}
	// 	fmt.Println(salaries)
}
