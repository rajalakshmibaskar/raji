package main

import "fmt"

func main() {
	var countries = []string{"india", "japan", "canada", "australia", "russia"}

	fmt.Printf("Countries: %v\n", countries)

	fmt.Printf(":2 %v\n", countries[:2])

	fmt.Printf("1:3 %v\n", countries[1:3])

	fmt.Printf("2: %v\n", countries[2:])

	fmt.Printf("2:5 %v\n", countries[2:5])

	fmt.Printf("0:3 %v\n", countries[0:3])
	
	fmt.Println(countries[:])
	
	fmt.Println(countries[0:])
	
	var employee = make(map[string]int)
	
	employee["velvizhi"] = 10
	
	employee["Shruthi"] = 20
	
	fmt.Println(employee)
 
	employeeList := make(map[string]int)
	
	employeeList["velvizhi"] = 10
	
	employeeList["Shruthi"] = 20
	
	fmt.Println(employeeList)
}