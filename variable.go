package main

import "fmt"

func main() {
	name := "Ajay"
	fmt.Println(name)

	name = "Zainul"
	fmt.Println(name)

	var (
		firstName = "Muhamad"
		midName   = "Zainul"
		lastName  = "kamal"
		nickName  = "Ajay"
	) // multiple variable
	fmt.Println(firstName, midName, lastName)
	fmt.Println(nickName)
}
