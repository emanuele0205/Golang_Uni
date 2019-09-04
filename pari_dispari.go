/*Problema: Scrivere un programma Go pari_dispari.go che,
dato un numero intero, determini se è pari o dispari.*/

package main 
import "fmt"
func main(){
	var n int
	fmt.Println("numero?")
	fmt.Scan(&n)
	if n%2==0{
		fmt.Println(n,"e' pari")
	}else{
		fmt.Println(n,"e' dispari")
	}

}