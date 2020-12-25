package main

import (
	"encoding/json"
	"fmt"
)

type Human struct {
	Name    string
	Age     int
	Address string
}

func main() {

	human1 := Human{"Raji", 23, "Delhi"}

	human_enc, err := json.Marshal(human1)

	if err != nil {

		fmt.Println(err)
	}

	fmt.Println(string(human_enc))

	human2 := []Human{
		{Name: "Raji", Age: 23, Address: "Delhi"},
		{Name: "Priya", Age: 21, Address: "Pune"},
		{Name: "Shiva", Age: 22, Address: "Bangalore"},
	}

	human2_enc, err := json.Marshal(human2)

	if err != nil {

		fmt.Println(err)
	}

	fmt.Println()
	fmt.Println(string(human2_enc))
}