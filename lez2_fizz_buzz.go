/*Problema Scrivere un programma Go lez2 fizz buzz.go che riceve in ingresso un numero intero e stampa ”Fizz” se il numero `e multiplo di 3, ”Buzz”
se il numero `e multiplo di 5, il numero stesso altrimenti.
Annotazioni Si osservi che le condizioni non sono mutuamente esclusive e
che quindi non si prestano a essere gestite con un if ...else if ....*/

package main 
import "fmt"
func main(){
	 var n int
	 fmt.Println("numero?")
	 fmt.Scan(&n)
	 if n%3==0{
		 if n%5==0{
			 fmt.Println("Fizz Buzz")
		 }else{
			 fmt.Println("Fizz")
		 }
	 }else if n%5==0{
		 fmt.Println("Buzz")
	 }else{
		 fmt.Println(n)
	 }
    
}