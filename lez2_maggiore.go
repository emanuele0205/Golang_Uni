/*Scrivere un programma Go lez2 maggiore.go che legge due interi e stampa il
maggiore.*/

package main
import "fmt"
func main(){
	var n1,n2 int
	fmt.Println("inserisci due numeri interi")
	fmt.Scan(&n1,&n2)
	if n1>n2 {
		fmt.Println(n1,"e' il maggiore")
	}else{
		fmt.Println(n2,"e' il maggiore")
	}
}