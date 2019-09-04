/*Scrivere un programma lez6_dado.go che analizza l'equita` 
di un dado contando la frequenza di apparizione dei suoi 
valori (1, 2, 3, 4, 5, 6). Il programma chiede all'utente 
quanti lanci fare. Utilizzare la funzione rand.Intn(n int)
del pacchetto "math/rand" e, per generare sequenze sempre 
diverse, importare il package "time" e usare l'istruzione 
rand.Seed(time.Now().Unix()) prima di iniziare a generare 
numeri.*/

package main 
import (
	"fmt"
	"math/rand"
	"time"
)
func main(){

	var n int 

	fmt.Println("quanti lanci di dado?")
	fmt.Scan(&n)
	 
	 var frequenza [7]int 


	rand.Seed(time.Now().Unix())
	for i:=0;i<n;i++{
		lancio:=rand.Intn(6)+1
		frequenza[lancio]++
	}

	for i:=1; i <= 6; i++{
		percent := float64(frequenza[i])*100/float64(n)
		fmt.Print(i, ": ", frequenza[i], " (", percent, "%)\n")
	}
}