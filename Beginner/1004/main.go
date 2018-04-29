/* Simple Product
Adapted by Neilor Tonin, URI  Brazil

Timelimit: 1
Read two integer values. After this, calculate the product between them and store the
result in a variable named PROD. Print the result like the example below. Do not
forget to print the end of line after the result, otherwise you will receive “Presentation Error”.

Input
The input file contains 2 integer numbers.

Output
Print PROD according to the following example, with a blank space before and after the equal signal.
*/
package main

import "fmt"

func main() {
	var a, b int
	fmt.Scanf("%d\n%d", &a, &b)
	PROD := a * b
	fmt.Printf("PROD = %d\n", PROD)
}
