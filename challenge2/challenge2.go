package main

import (
	"fmt"
)

func main() {

	fmt.Println("****** Challenge 2 ******")
	orderedList := []float64{2, 2, 8, 10}
	challenge2(orderedList)

}

func challenge2(list []float64) {
	//calculating the mean
	mean := mean(list)
	fmt.Println("Mean -> ", mean)

	//calculating Median
	median := median(list)
	fmt.Println("Median -> ", median)

	//calculating maxValue
	maxValue := maxValue(list)
	fmt.Println("Max value -> ", maxValue)

}

func maxValue(list []float64) float64 {
	return list[len(list)-1]
}

func mean(list []float64) float64 {
	mean := 0.0
	for _, num := range list {
		mean += num
	}
	mean /= float64(len(list))
	return mean
}

func median(list []float64) float64 {

	median := 0.0
	if len(list)%2 == 1 {
		median = list[len(list)/2+1]
	} else {
		median = (list[len(list)/2-1] + list[len(list)/2]) / 2
	}

	return median
}
