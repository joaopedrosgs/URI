package main

import (
	"fmt"
	"math"
)

const pi = 3.14159

func main() {
	var r float64
	fmt.Scanf("%f", &r)
	fmt.Printf("A=%.4f\n", math.Pow(r, 2)*pi)
}
