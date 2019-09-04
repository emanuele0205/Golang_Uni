/*Scrivere un programma lez3_min_max.go che legge
da tastiera una serie di numeri (n numeri? o != 0 
fino a incontrare 0) e stampa il massimo e il minimo
(ricerca del valore estremo, most-wanted value)*/

package main
import "fmt"
func main(){

	

	var min, max, n int
	fmt.Println("inserire i numeri da calcolare max e min, premere 0 per uscire")
	fmt.Scan(&n)
	min=n
	for n!=0{
		fmt.Scan(&n)
		if n==0{
			break
		}
		if n>max{
			max=n
		}else if n<min{
			min=n
		}

		}
	
	fmt.Println("max:",max)
	fmt.Println("min",min)
}






