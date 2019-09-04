/*Scrivere un programma lez4_v.go che legge un intero 
positivo n e stampa una lettera 'V' di asterischi di 
altezza n.
*/

package main 
import "fmt"
func main(){

	var n,h int 
	fmt.Println("inserire altezza n")
	fmt.Scan(&n)
h=n
for j:=1;j<n;j++{
	for t := 0; t < j; t++ {
		fmt.Print(" ")
	}
	fmt.Print("*")
	for t:=0;t<h;t++{
		fmt.Print(" ")
	}
	fmt.Println("*")
	h=h-2
}
for k:=0;k<n;k++{
	fmt.Print(" ")
}
fmt.Println("*")
}