/* Banknotes
Adapted by Neilor Tonin, URI  Brazil

Timelimit: 1
In this problem you have to read an integer value and
calculate the smallest possible number of banknotes in
which the value may be decomposed. The possible
banknotes are 100, 50, 20, 10, 5, 2 e 1. Print the
read value and the list of banknotes.

Input
The input file contains an integer value N (0 < N < 1000000).

Output
Print the read number and the minimum quantity of
each necessary banknotes in Portuguese language, as
the given example. Do not forget to print the end of
line after each line, otherwise you will receive
“Presentation Error”.
*/
package main

import "fmt"

/*
// Faz exatamente a mesma coisa, mas o URI não quer aceitar
import (
	"fmt"
	"sort"
)

func Notas(valor int, notasDisponiveis [] int) map[int]int {
	sort.Ints(notasDisponiveis)
	notasNecessarias := make(map[int]int)
	for i := len(notasDisponiveis) - 1; i >= 0; i-- {
		notasNecessarias[notasDisponiveis[i]] = valor / notasDisponiveis[i]
		valor = valor % notasDisponiveis[i]
	}
	return notasNecessarias
}
func main() {

	var valor int
	fmt.Scanf("%d", &valor)
	if valor <= 0 || valor >= 1000000 {
		return
	}
	println(valor)
	notasExistentes := []int{100, 50, 20, 10, 5, 2, 1}
	notasNecessarias := Notas(valor, notasExistentes)
	for i := len(notasExistentes) - 1; i >= 0; i-- {
		fmt.Printf("%d nota(s) de R$ %d,00\n", notasNecessarias[notasExistentes[i]], notasExistentes[i])
	}
}
*/

func main() {

	var valor int
	fmt.Scanf("%d", &valor)
	fmt.Printf("%d\n", valor)
	fmt.Printf("%d nota(s) de R$ 100,00\n", valor/100)
	valor %= 100
	fmt.Printf("%d nota(s) de R$ 50,00\n", valor/50)
	valor %= 50
	fmt.Printf("%d nota(s) de R$ 20,00\n", valor/20)
	valor %= 20
	fmt.Printf("%d nota(s) de R$ 10,00\n", valor/10)
	valor %= 10
	fmt.Printf("%d nota(s) de R$ 5,00\n", valor/5)
	valor %= 5
	fmt.Printf("%d nota(s) de R$ 2,00\n", valor/2)
	valor %= 2
	fmt.Printf("%d nota(s) de R$ 1,00\n", valor)
}
