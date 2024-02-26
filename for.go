package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		fmt.Println("Perulangan ke-", i)
	}
	fmt.Println("Done")

	names := []string{"Muhamad", "Zainul", "Kamal"}
	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

	for index, name := range names {
		fmt.Println("index", index, "=", name)
	}
}
