package main

import "fmt"

func main() {
	var a [5]int
	fmt.Println(a)

	b := [4]int{1, 2, 3, 4}
	fmt.Println(b)

	c := [...]int{1, 2, 3}
	fmt.Println(c)

	d := [...]int{100, 4: 400, 500, 600}
	fmt.Println(d)

	a2 := [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	fmt.Println(a2)

	var a3 [2][4]int
	for i := range 2 {
		for j := range 4 {
			a3[i][j] = i + j
		}
	}
	fmt.Println(a3)
}
