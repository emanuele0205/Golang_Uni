/*Problema: Scrivere un programma Go quoziente_resto.go che,
dati un dividendo e un divisore (interi), calcoli il quoziente e il
resto.

Annotazioni L’operatore per la divisione (/) tra interi calcola la
parte intera del risultato. L’operatore per il resto della divisione
è %.*/

package main

import "fmt"

func main() {
	var d1, d2 int
	var quoziente, resto float64

	fmt.Println("inserire dividendo e divisore interi")
	fmt.Scan(&d1, &d2)

	quoziente = float64(d1 / d2)
	resto = float64(d1 % d2)

	fmt.Println("quoziente: ", quoziente)
	fmt.Println("resto: ", resto)

}
