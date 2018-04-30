/* Distance Between Two Points
Adapted by Neilor Tonin, URI  Brazil

Timelimit: 1
Read the four values corresponding to the x and y axes of two
points in the plane, p1 (x1, y1) and p2 (x2, y2) and calculate
the distance between them, showing four decimal places after
the comma, according to the formula:

Distance =

Input
The input file contains two lines of data. The first one
contains two double values: x1 y1 and the second one also
contains two double values with one digit after the decimal
point: x2 y2.

Output
Calculate and print the distance value using the provided
formula, with 4 digits after the decimal point.
*/
package main

import (
	"fmt"
	"math"
)

func main() {
	var x1, y1, x2, y2, distance float64

	fmt.Scanf("%f	 %f\n%f %f", &x1, &y1, &x2, &y2)
	distance = math.Sqrt(math.Pow(x2-x1, 2) + math.Pow(y2-y1, 2))

	fmt.Printf("%.4f\n", distance)
}
