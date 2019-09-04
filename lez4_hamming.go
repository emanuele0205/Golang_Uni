/*
La distanza di Hamming tra due stringhe e` data dal 
numero di caratteri che sono diversi dai corrispondenti 
(cioe` nella stessa posizione) nell'altra stringa.

Scrivere un programma lez4_hamming.go che legge due
parole e calcola la distanza di Hamming tra loro.

Esempio
-------
Scrivi due parole:
rosa roba
1*/

package main 
import "fmt"
func main(){

	var(
		s1,s2 string
		cont int 
	)

	fmt.Println("inserire le due stringhe ")
	fmt.Scan(&s1,&s2)
	for j:=0;j<len(s1);j++{
		if s1[j]!=s2[j]{
		cont++
		}
	}
	fmt.Printf("distanza di hamming %d\n",cont)
}