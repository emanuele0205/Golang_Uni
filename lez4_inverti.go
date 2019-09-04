/*Scrivere un programma lez4_inverti.go che legge una 
stringa di testo e la visualizza in ordine inverso 
(dall'ultimo carattere al primo).
*/

package main 
import "fmt"
func main(){

	var	 s string

fmt.Println("inserire stringa ")
fmt.Scan(&s)
t:=len(s)
s2:=make([]rune, len(s))
for _,char:=range(s){
	s2[t-1]+=char
	t--
}
for _,char2:=range s2{
	fmt.Printf("%c",char2)
}

fmt.Println()
}