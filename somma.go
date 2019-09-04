/*Problema: Scrivere un programma Go somma.go che, dati tre
numeri interi, calcoli la loro somma.*/

package main

import "fmt"

func main() {
	var n1, n2, n3 float64
	fmt.Println("inserire i tre numeri da sommare")
	fmt.Scan(&n1, &n2, &n3)

	somma := n1 + n2 + n3
	fmt.Println(n1, " + ", n2, " + ", n3, " = ", somma)

}
