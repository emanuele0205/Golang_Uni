/*Consideriamo la seguente tabella che associa uno sconto a una persona a seconda
della sua et`a
et`a sconto
0 - 9 50%
10 - 16 30%
17 - 25 20%
26 - 45 30%
46 - 59 20%
60 - 75 30%
76 . . . .... 50%
Scrivere un programma Go lez2 sconti.go che legge un intero n >= 0,
rappresentante l’et`a di una persona, e stampa lo sconto associato. Si assuma
che il numero inserito sia maggiore o uguale a 0 (quindi non `e necessario fare il
controllo n >= 0).
Annotazioni Si noti che gli sconti possono assumere solamente i valori 20, 30
e 50. Quindi il codice deve avere la struttura:
leggi n
if(CONDIZIONE_1)
stampa 20
else if( CONDIZIONE_2)
stampa 50
else
stampa 30
dove CONDIZIONE 1 e CONDIZIONE 2 vanno opportunamente definite
usando gli operatori booleani.*/

package main 
import "fmt"
func main(){
	var eta int 
	fmt.Println("Eta':")
	fmt.Scan(&eta)
	if eta>=17 && eta<=25 || eta>=46 && eta<=59{
		fmt.Println("sconto: 20%")
	}else if(eta>=0 && eta<=9 || eta>=76){
		fmt.Println("sconto: 50%")
	}else{
		fmt.Println("sconto: 30%")
	}
}