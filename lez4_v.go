/*Scrivere un programma lez4_v.go che legge un intero 
positivo n e stampa una lettera 'V' di asterischi di 
altezza n.
*/

package main

import (
	"fmt"
)

func main() {
	var n int
	s := " "

	fmt.Scan(&n)
	for i := 1; i < n; i++ {
		fmt.Print(s, "*")
		for j := i; j < n; j++ {
			if j == n-1 {
				fmt.Print(" ")
			} else {
				fmt.Print("  ")
			}

		}

		fmt.Print("*")
		fmt.Println()
		s += " "
	}
	for z := 0; z <= n; z++ {
		if z == n {
			fmt.Println("*")
		} else {
			fmt.Print(" ")
		}

	}

}
