package main

import "fmt"

func main() {
	n := 10

	if n%2 == 0 {
		fmt.Println("Number is even")
	} else {
		fmt.Println("Number is odd")
	}

	if n <= 9 {
		fmt.Println("It's a digit")
	} else if n == 10 {
		fmt.Println("It's dec")
	} else {
		fmt.Println("It's numeric")
	}
}
