package main

import "fmt"

type Address struct {
	Name    string
	city    string
	Pincode int
}

func main() {

	var a Address
	fmt.Println(a)

	a1 := Address{"Meenakshi", "Kerala", 689500}

	fmt.Println("Address1: ", a1)

	a2 := Address{Name: "Rajalakshmi", city: "Ranipet",
		Pincode: 632401}

	fmt.Println("Address2: ", a2)

	a3 := Address{Name: "Praveen kumar"}

	fmt.Println("Address3: ", a3)
}