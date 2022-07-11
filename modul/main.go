package main

import "fmt"
import "os"

func main() {
	var n int
	fmt.Print("Enter your number: ")
	fmt.Fscan(os.Stdin, &n)
	
	if n < 0 {
		n = -n
	}
	fmt.Println("Modul", n)
}