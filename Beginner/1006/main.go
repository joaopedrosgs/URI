/* Average 2
Adapted by Neilor Tonin, URI  Brazil

Timelimit: 1
Read three values (variables A, B and C), which are the three student's grades.
Then, calculate the average, considering that grade A has weight 2, grade B
has weight 3 and the grade C has weight 5. Consider that each grade can go
from 0 to 10.0, always with one decimal place.

Input
The input file contains 3 values of floating points with one digit after the decimal point.

Output
Print MEDIA(average in Portuguese) according to the following example,
with a blank space before and after the equal signal.
*/
package main

import "fmt"

const (
	pesoA = 2
	pesoB = 3
	pesoC = 5
)

func main() {
	var a, b, c float64
	fmt.Scanf("%f\n%f\n%f", &a, &b, &c)
	var MEDIA float64 = ((a * pesoA) + (b * pesoB) + (c * pesoC)) / (pesoA + pesoB + pesoC)
	fmt.Printf("MEDIA = %.1f\n", MEDIA)
}
