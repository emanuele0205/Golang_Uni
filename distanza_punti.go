/*Problema: Scrivere un programma Go distanza_punti.go che
calcola la distanza tra due punti nel piano cartesiano.*/

package main 
import (
	"fmt"
	"math"
)
func main(){

	var x1,x2,y1,y2,distanza,p1,p2,xtot,ytot float64
	fmt.Println("inserisci la x e la y del primo punto")
	fmt.Scan(&x1,&y1)
	fmt.Println("inserisci la x e la y del secondo punto")
	fmt.Scan(&x2,&y2)
	xtot=x2-x1
	ytot=y2-y1
	p1=math.Pow(xtot, 2)
	p2=math.Pow(ytot, 2)
	distanza=math.Sqrt(p1+p2)
	fmt.Println("La distanza tra i due punti e': ",distanza)
}