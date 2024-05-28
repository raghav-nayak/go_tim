package main

import (
	"fmt"
	"math"
)

func main() {
	var num1 float64 = 8
	var num2 int = 3
	answer := num1 / float64(num2)
	fmt.Printf("%f\n", answer)

	num3 := 9
	num4 := 4
	answer2 := float64(num3 / num4)
	fmt.Printf("%g\n", answer2)

	var num5 int = 4
	var num6 int = 4
	answer3 := num5 / num6
	fmt.Printf("%d\n", answer3)

	fmt.Printf("%f\n", math.Ceil(3.4))
}