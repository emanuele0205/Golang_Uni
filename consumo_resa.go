/*Problema. Scrivere un programma Go consumo_resa.go che
calcola il consumo medio e la resa di un motore data la distanza
totale percorsa (in km) e la quantità di carburante utilizzata (in
litri). I valori sono di tipo float64.
Annotazioni.. Il consumo medio di carburante si esprime in l/km
ed è la quantità di carburante che occorre in media per
percorrere un km di strada. La resa di un motore è data dalla
distanza percorsa in media con un litro di carburante e si
espreime in km/l.*/

package main 
import "fmt"
func main(){

	var cmedio,resa,distot,carburante float64
	fmt.Println("inserisci la distanza percorsa in km")
	fmt.Scan(&distot)
	fmt.Println("inserisci la quantità di carburante utilizzata in L")
	fmt.Scan(&carburante)

	cmedio=carburante/distot
	resa=distot/carburante

	fmt.Println("consumo medio: ",cmedio," l/km")
	fmt.Println("resa media: ",resa," km/l")

}