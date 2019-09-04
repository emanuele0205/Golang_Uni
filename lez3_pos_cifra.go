/*Scrivere un programma lez3_pos_cifra.go che legge da tastiera 
un numero intero n di quattro cifre e un numero i di una cifra e 
stampa in che posizione (1 per le unita`) e` la prima occorrenza 
di i nella sequenza di cifre che rappresentano n 
(0 se i non e` presente).*/

package main
import "fmt"
func main(){

	var n,i,flag int
	fmt.Println("inserire un numero di 4 cifre e un successivo numero da una cifra ")
	fmt.Scan(&n,&i)
	for t:=1;t<=n;t++{
		if n%10==i{
			flag=t
			break
		}else{
			n=n/10
		}
	}
fmt.Println("l'occorrenza sta nella posizione: ",flag)
}