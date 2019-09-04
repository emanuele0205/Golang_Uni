/*Scrivere un programma Go lez2 voto valido.go che legge un numero intero.
Se il numero non `e compreso tra 0 e 30, stampa “voto non valido”.*/

package main 
import "fmt"
func main(){
	var n int 
	fmt.Println("voto?")
	fmt.Scan(&n)
	if(n<0||n>30){
		fmt.Println("valore non valido")
	}
}