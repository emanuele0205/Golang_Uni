/*Problema: Scrivere un programma Go operazioni.go che, dati
due numeri float64, esegua addizione, sorazione,
moltiplicazione e divisione dei due numeri.*/

package main 
import "fmt"
func main(){
	var n1,n2,add,sott,molt,div float64
	fmt.Println("inserire due numeri")
	fmt.Scan(&n1,&n2)
	add=n1+n2
	sott=n1-n2
	molt=n1*n2
	div=n1/n2
	fmt.Println("addizione: ",add)
	fmt.Println("sottrazione: ",sott)
	fmt.Println("moltiplicazione: ",molt)
	fmt.Println("divisione: ",div)
}