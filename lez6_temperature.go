/*Scrivere un programma lez6_temperature.go che legge 10 temperature 
(int) da tastiera e calcola 
- la media
- la mediana
- il numero di temperature sotto la media delle temperature stesse.
*/

package main 
import(
		"fmt"
		"sort"
)
func main(){
	const dim = 10
	var temperatura, copiaDiTemperatura [dim]int
	var somma int
	var media float64

	fmt.Println("Scrivi", dim, "valori int (temperature)")
	for i:= 0; i < dim; i++{
		fmt.Scan(&temperatura[i])
	}

	fmt.Println("temperature:", temperatura)

	for i:= 0; i < dim; i++{
		somma += temperatura[i]
	}

	media = float64(somma)/dim

	fmt.Println("media:", media)

	copiaDiTemperatura = temperatura

	fmt.Println("mediana:", mediana(copiaDiTemperatura[:], dim))

	var count int
	for i:= 0; i < dim; i++{
		if float64(temperatura[i]) < media{
			count++
		}
	}
	fmt.Println("n. di temperature sotto la media:", count)
}

func mediana(a []int, size int)  (med float64){
	sort.Ints(a)
	mid := int(size / 2)
	if size %2 == 0{
		med = float64(a[mid-1] + a[mid]) / 2
	} else{
		med = float64(a[mid])
	}
	return
}