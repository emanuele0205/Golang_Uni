/*Problema: Scrivere un programma Go
centigradi_fahrenheit.go che converta in Fahrenheit una
temperatura espressa in gradi centigradi (float64).
Annotazioni. Un grado Fahrenheit è 5/9 di un grado centigrado e
0 gradi centigradi corrisponde a 32 gradi Fahrenheit.*/

package main 
import "fmt"
func main(){
	var cent,far float64

	fmt.Println("inserire i gradi espressi in centigradi")
	fmt.Scan(&cent)
	far=((cent/5)*9)+32

	fmt.Println(cent," C = ",far," F")
}