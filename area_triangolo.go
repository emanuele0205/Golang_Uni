/*Problema: Scrivere un programma Go area_triangolo.go che,
date le ampiezze di due angoli di un triangolo, determini
l’ampiezza del terzo angolo.*/

package main 
import "fmt"
func main(){
	var a1,a2,a3 float64

	fmt.Println("inserire ampiezza di due angoli di un triangolo")
	fmt.Scan(&a1,&a2)

	a3=180-a1-a2
	
	fmt.Println("ampiezza terzo angolo: ",a3)

}