package main

import (
	"fmt"
	"math"
	"os"
)

func main() {
	var x float64
	fmt.Print("Enter angle (degrees): ")
	fmt.Fscan(os.Stdin, &x)
	r := x * math.Pi / 180
	sin, cos := math.Sincos(r)
	tg := sin / cos
	ctg := cos / sin

	fmt.Printf("Sin = %.1f Cos = %.1f\n Tg = %.1f Ctg = %.1f", sin, cos, tg, ctg)

}

