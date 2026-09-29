package main

import "fmt"

func main() {
	for {
		fmt.Println("It's infinite loop...")
		fmt.Println("Breaking...")
		break
	}
	fmt.Println()

	fmt.Println("Simple 'while' loop with single condition:")
	i := 0
	for i <= 3 {
		fmt.Println(i)
		i++
	}
	fmt.Println()

	fmt.Println("Basic for with intial/condition/after:")
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
	fmt.Println()

	fmt.Println("Ranged for:")
	for i := range 4 {
		fmt.Println(i)
	}
}
