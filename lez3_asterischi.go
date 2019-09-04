/*Scrivere un programma lez3_asterischi.go che, dato un intero n, stampa n asterischi.
(ripetizione, stepper)
*/

package main 
import "fmt"
func main(){
	var n int
	fmt.Println("inserisci il numero di asterischi")
	fmt.Scan(&n)
	for i:=0;i<n;i++{
		fmt.Print("*")
	}
	fmt.Println()

}