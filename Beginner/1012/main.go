/* Area
Adapted by Neilor Tonin, URI  Brazil

Timelimit: 1
Make a program that reads three floating point values: A, B and C.
Then, calculate and show:
a) the area of the rectangled triangle that has base A and height C.
b) the area of the radius's circle C. (pi = 3.14159)
c) the area of the trapezium which has A and B by base, and C by height.
d) the area of ​​the square that has side B.
e) the area of the rectangle that has sides A and B.

Input
The input file contains three double values with one digit after the decimal point.

Output
The output file must contain 5 lines of data. Each line corresponds to
one of the areas described above, always with a corresponding message
(in Portuguese) and one space between the two points and the value.
The value calculated must be presented with 3 digits after the decimal point.
*/
package main

import (
	"fmt"
	"math"
)

const pi = 3.14159

func main() {
	var a, b, c, triangle, circle_volume, trapezium_area, square_area, rectangle_area float64

	fmt.Scanf("%f %f %f", &a, &b, &c)
	triangle = a * c / 2
	circle_volume = pi * math.Pow(c, 2)
	trapezium_area = (a + b) * c / 2
	square_area = b * b
	rectangle_area = a * b

	fmt.Printf("TRIANGULO: %.3f\n", triangle)
	fmt.Printf("CIRCULO: %.3f\n", circle_volume)
	fmt.Printf("TRAPEZIO: %.3f\n", trapezium_area)
	fmt.Printf("QUADRADO: %.3f\n", square_area)
	fmt.Printf("RETANGULO: %.3f\n", rectangle_area)
}
