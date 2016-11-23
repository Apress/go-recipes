package main

import (
	"encoding/json"
	"fmt"
)

// Employee struct
type Employee struct {
	ID                            int
	FirstName, LastName, JobTitle string
}

func main() {
	emp := Employee{
		ID:        100,
		FirstName: "Shiju",
		LastName:  "Varghese",
		JobTitle:  "Architect",
	}
	// Encoding to JSON
	data, err := json.Marshal(emp)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	jsonStr := string(data)
	fmt.Println("The JSON data is:")
	fmt.Println(jsonStr)

	b := []byte(`{"ID":101,"FirstName":"Irene","LastName":"Rose","JobTitle":"Developer"}`)
	var emp1 Employee
	// Decoding JSON data to a value of struct type
	err = json.Unmarshal(b, &emp1)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("The Employee value is:")
	fmt.Printf("ID:%d, Name:%s %s, JobTitle:%s", emp1.ID, emp1.FirstName, emp1.LastName, emp1.JobTitle)
}
