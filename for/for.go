package main

import "fmt"

func main() {
	multiply := 1
	for i := 1; i < 11; i++ {
		multiply *= i
	}
	subtraction := 1000
	for i := 1; i < 11; i++ {
		subtraction -= i
	}
	addition := 0
	for i := 1; i < 102; i = i + 2 {
		addition += i
	}
	division := 1000000.0
	for i := 2.0; i < 11; i++ {
		division /= i
	}
	fmt.Println(multiply, subtraction, addition, division)
}
