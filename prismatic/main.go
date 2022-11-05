package main

import (
	"fmt"
	"math"
)

func main() {
	total := 4
	for i := -4; i <= total; i++ {
		a := total - int(math.Abs(float64(i)))
		for k := 1; k <= total-int(a); k++ {
			fmt.Print(" ")
		}

		for j := 1; j <= 2*a-1; j++ {
			fmt.Print("*")

		}
		fmt.Println()
	}

}
