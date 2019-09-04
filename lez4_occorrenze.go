/*Scrivere un programma lez4_occorrenze.go che legge 
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

	var (
		s string 
		r string 
		cont int
	)
	fmt.Println("inserire la stringa e il carattere")
	fmt.Scan(&s)
	fmt.Scan(&r)
	t:=[]rune(s)
	for _,char:=range(t){
		if string(char)==r{
			cont++
		
		}
	}
	fmt.Println(cont)
	
}