package main

import "fmt"

func main() {
	i := 2

	switch i {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	}

	switch {
	case i > 5:
		fmt.Println("Variable is too big!")
	default:
		fmt.Printf("Variable: %d\n", i)
	}

	whatType := func(i any) {
		switch t := i.(type) {
		case int:
			fmt.Println("This is int")
		case float32, float64:
			fmt.Println("This is float")
		default:
			fmt.Printf("Unknown type '%T'\n", t)
		}
	}

	whatType(64)
	whatType(34.6)
	whatType("string")
}
