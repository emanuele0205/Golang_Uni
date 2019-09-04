/*Scrivere un programma lez3_euclide.go che, dati due interi,
 calcolare il MCD tra i due numeri con l'algoritmo di Euclide. */

package main 
import "fmt" 
func main(){
	var a,b int
	fmt.Println("inserisci i due numeri con cui calcolare mcd")
	fmt.Scan(&a,&b)
	if b==0{
		fmt.Println("mcd tra i due numeri e':",a)
	}else{
		for b!=0{
			r:=a%b
			a=b
			b=r
		}
	fmt.Println("mcd tra i due numeri e':",a)
	}
	
}