/*Problema: Scrivere un programma Go kmora_migliaora.go che
converte kilometri all’ora (km/h) in miglia all’ora (mi/h). Un
miglio equivale a 1.609344 km.
Indicazione. Usare una costante MI_TO_KM = 1.609344.*/

package main 
import "fmt"
func main(){
	const MI_TO_KM=1.609344
	var km,mi float64

	fmt.Println("inserisci la misura della velocita' in km/h")
	fmt.Scan(&km)
	mi=km/MI_TO_KM
	fmt.Println(km," km/h = ",mi," mi/h")
}