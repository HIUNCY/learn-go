package main

import "fmt"

func main() {
	for i := 0; i <= 10; i++ {
		if i == 5 {
			break // stop the loop when i = 5
		}
		fmt.Println("Perulangan ke-", i)
	}

	for i := 0; i <= 10; i++ {
		if i == 5 {
			continue // stop the loop when i = 5 and continue the next iteration
		}
		fmt.Println("Perulangan ke-", i)
	}
}
