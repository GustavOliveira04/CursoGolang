package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	jsonString := `{"name":"Gustavo", "age":21}`
	var person Person
	json.Unmarshal([]byte(jsonString), &person)
	fmt.Printf("Nome: %s, Idade: %d\n", person.Name, person.Age)

	person.Name = "Carlos"
	person.Age = 52

	jsonData, _ := json.Marshal(person)
	fmt.Println(string(jsonData))
}
