/* The Greatest
Adapted by Neilor Tonin, URI  Brazil

Timelimit: 1
Make a program that reads 3 integer values and present the
greatest one followed by the message "eh o maior". Use the following formula:

maiorAB=(A+B+ABS(A-B))/2

Input
The input file contains 3 integer values.

Output
Print the greatest of these three values followed by a space and the message “eh o maior”.
*/
package main

import (
	"fmt"
	"math"
)

func maiorAB(a, b int) int {
	return (a + b + int(math.Abs(float64(a)-float64(b)))) / 2
}

func main() {
	var a, b, c, maior int

	fmt.Scanf("%d %d %d", &a, &b, &c)
	maior = maiorAB(a, b)
	maior = maiorAB(maior, c)

	fmt.Printf("%d eh o maior\n", maior)
}
