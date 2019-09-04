/*Problema: Scrivere un programma Go cerchio.go che, data la
misura del raggio di un cerchio, calcoli la circonferenza e l’area.
Annotazioni La circonferenza di un cerchio di raggio r è data da
2πr.
L’area delimitata da un cerchio di raggio r è data da πr
2
.
La leera greca π rappresenta una constante di valore circa
uguale a 3.14.
Indicazione. Usare una costante Pi = 3.14.*/

package main
import "fmt"
func main(){
	const Pi=3.14
	var r,circonferenza,area float64

	fmt.Println("inserisci la misura del raggio (in cm)")
	fmt.Scan(&r)
	circonferenza=2*Pi*r
	area=Pi*r*r
	fmt.Println("Circonferenza = ",circonferenza," cm")
	fmt.Println("Area = ",area," cm^ 2")

}