/* Difference
Adapted by Neilor Tonin, URI  Brazil

Timelimit: 1
Read four integer values named A, B, C and D. Calculate and print
the difference of product A and B by the product of C and D (A * B - C * D).

Input
The input file contains 4 integer values.

Output
Print DIFERENCA (DIFFERENCE in Portuguese) with all the capital letters,
according to the following example, with a blank space before and after
the equal signal.
*/
package main

import "fmt"

func main() {
	var a, b, c, d int
	fmt.Scanf("%d\n%d\n%d\n%d", &a, &b, &c, &d)
	DIFERENCA := (a * b) - (c * d)
	fmt.Printf("DIFERENCA = %d\n", DIFERENCA)
}
