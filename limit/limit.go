package main

import (
	"fmt"
	"os"
)

func main() {

	var money float64

	fmt.Print("Your investment in roubles: ")
	fmt.Fscan(os.Stdin, &money)

	if money < 0 {
		fmt.Println("Please, enter positive sum")

		os.Exit(1)
	}

	if money > 1400000 {
		fmt.Println("Sorry, you have too much money to be insured")

		os.Exit(2)
	}

	const rate = 1.05
	const limit = 1400000.0

	years := 0

	for money < limit {
		money *= rate
		years++
	}

	years--
	money /= rate

	fmt.Printf("Years insured = %v\n Maximum money insured = %.1f\n", years, money)
}
