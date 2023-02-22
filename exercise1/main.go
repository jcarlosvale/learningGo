package main

import "fmt"

func main() {
	numbers := []int{}
	for i := 0; i < 11; i++ {
		numbers = append(numbers, i)
	}

	for _, n := range numbers {
		if n%2 == 0 {
			fmt.Println(n, " is even")
		} else {
			fmt.Println(n, " is odd")
		}
	}
}
