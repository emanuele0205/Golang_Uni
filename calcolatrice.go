/*Calcolatrice*/

package main
import "fmt"
func main(){

	var n1,n2 float64
	var s string 
	fmt.Println("inserisci i due numeri")
	fmt.Scan(&n1,&n2)
	fmt.Println("scegli l'operazione: somma,sottrazione,moltiplicazione,divisione")
	fmt.Scan(&s)
	switch  {
	case s=="somma":
		fmt.Println(n1+n2)
	case s=="sottrazione":
		fmt.Println(n1-n2)
	case s=="moltiplicazione":
		fmt.Println(n1*n2)
	case s=="divisione":
		fmt.Println(n1/n2)
	default:
		break
	}


}