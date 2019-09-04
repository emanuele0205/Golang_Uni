/*Problema: Scrivere un programma Go minuti.go che calcola il
numero totale di minuti, dati in input il numero di ore e di
minuti.
Indicazione. Usare una costante MIN_IN_HOUR = 60.*/

package main 
import "fmt"
func main(){
	const MIN_IN_HOUR=60
	var ore, minuti,tot int 

	fmt.Println("inserisci ore e minuti")
	fmt.Scan(&ore,&minuti)
	tot= ore*MIN_IN_HOUR+minuti
	fmt.Println("minuti totali = ",tot)

}