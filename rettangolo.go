/*Problema: Scrivere un programma Go rettangolo.go che, date
le misure dell’altrezza e della base di un reangolo (int), calcoli il
perimetro e l’area.*/

package main
import "fmt"
func main(){
	var altezza,base,area,perimetro int

	fmt.Println("inserisci base e altezza (in cm) separate da uno spazio")
	fmt.Scan(&base,&altezza)
	area=base*altezza
	perimetro=2*(base+altezza)
	
	fmt.Println("il perimetro del rettangolo e' ",perimetro," cm")
	fmt.Println("l'area del rettangolo e' ",area," cm^ 2")
}