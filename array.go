package main

import "fmt"

func main() {
	names := [3]string{
		"Muhamad",
		"Ajay",
		"Kamal",
	} // you can replace [3] to [...] if you want push element to array more than 3
	fmt.Println(names)
	names[1] = "Zainul" // Change value
	fmt.Println(names)
	fmt.Println(len(names)) // len: length of array
}
