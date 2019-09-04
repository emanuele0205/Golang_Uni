/*Scrivere un programma lez3_andamento.go che 
legge da tastiera una serie di numeri > -1 e stampare 
"andamento positivo" ogni volta che il nuovo valore e` maggiore 
o uguale al precedente e "andamento negativo" altrimenti. 
Si ferma quando il numero in input e` -1.*/

package main 
import "fmt"
func main(){

	var n,flag int
	fmt.Println("inserire numeri per uscire inserire -1")
	fmt.Scan(&n)
	flag=n
	for n!=-1{
		fmt.Scan(&n)
		if n==-1{
			break
		}
		if n==flag || n>flag{
			fmt.Println("andamento positivo")
			flag=n
		}else{
			fmt.Println("andamento negativo")
			flag=n
		}
	}

}