package main

import "fmt"

func main() {
	person := map[string]string{
		"name":    "Zainul",
		"address": "Jakarta",
	}
	fmt.Println(person)
	person["name"] = "Ajay" // change the value
	fmt.Println(person)
	fmt.Println(len(person))  // len: length of map
	delete(person, "address") // delete value
	fmt.Println(person)
}
