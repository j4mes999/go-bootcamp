package main

import (
	"fmt"
	"math"
	"math/rand"
)

const N = 90000000

func main() {

	fmt.Println("****** Challenge 3 ******")
	challenge3()

}

//Montecarlo Simulation
func challenge3() {

	sum := 0.0
	for i := 0; i < N; i++ {
		x1 := rand.Float64()
		x2 := rand.Float64()
		z := math.Sqrt(math.Pow(x1, 2) + math.Pow(x2, 2))
		if z <= 1 {
			sum++
		}

	}

	res := 4 * (sum / N)
	fmt.Println("Pi approximation using Monte Carlo method is -> ", res)
}
