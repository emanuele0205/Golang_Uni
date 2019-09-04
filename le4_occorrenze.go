/*
Scrivere un programma lez4_occorrenze.go che legge 
una stringa e un carattere e stampa quante occorrenze 
di quel carattere ci sono nella stringa

Esempio
-------

Stringa e carattere:
giornata a
2*/

package main 
import "fmt"
func main(){

	var s string
	var r rune 
	var count int 
	fmt.Println("inserisci la frase e il carattere ")
	fmt.Scan(&s,&r)
	for _,c:= range s{
		if r == c{
			fmt.Println(c,r,count)
		count++
	} 
		}
fmt.Println("occorrenze: ",count)
}