/*
Scrivere un programma lez4_bin_to_int.go che legge 
una stringa composta di soli 0 e 1 e fa la conversione 
da binario a decimale del numero binario rappresentato 
dalla stringa e stampa il risultato.
*/

package main
import (
		"fmt"
		"math"
)
func main(){

	var ( 
		s string 
		somma float64
	)	
	fmt.Println("inserire numero binario")
	fmt.Scan(&s)
	t:=float64(len(s))
	for i:=0;i<len(s);i++{
		if s[i]=='1'{
			somma+=math.Pow(2,t-1)
		}
		t--
		}
		fmt.Println(somma)
	}
