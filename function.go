package main

import "fmt"

func sayHello(name string) string {
	return "Hello, " + name

}

func getFullName() (string, string, string) {
	return "Muhamad", "Zainul", "Kamal" // return multiple values
}

func sumAll(numbers ...int) int { // "..." means numbers will receive more than one value
	total := 0

	for i := 0; i < len(numbers); i++ {
		total += numbers[i]
	}
	return total
}

func main() {
	fmt.Println(sayHello("Ajay"))

	firstName, midName, lastName := getFullName() // u can put "_" in variable if u dont need the value
	fmt.Println(firstName, midName, lastName)

	fmt.Println(sumAll(10, 5, 10, 10))

	numbers := []int{10, 30, 5}
	fmt.Println(sumAll(numbers...)) // in case u already have slice, u can put "..." after variable
}
