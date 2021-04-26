package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) (float64, int) {
	z := x
	i := 1
	for ; i < 10; i++ {

		z -= (z*z - x) / (2 * z)

		if z-math.Sqrt(x) == 0 {
			break
		}

	}

	return z, i
}

func main() {
	value, iteration := Sqrt(2)
	fmt.Println("The value of z is ", value, " and it matched at iteration", iteration, "with math.Sqrt")
}
