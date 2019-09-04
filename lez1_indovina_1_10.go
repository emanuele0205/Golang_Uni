/*Scrivere un programma Go lez2 indovina 1 10.go che fissa un numero intero
tra 1 e 10 da indovinare, legge un intero da standard input e,
• se il numero in input e fuori dall’intervallo 1-10, stampa “Valore non
valido”;
• se il numero `e quello fissato, stampa “Hai indovinato!”;
• altrimenti stampa “Ciao”.*/

package main 
import "fmt"
func main(){
	const nfisso=2
	var n int 
	fmt.Println("inserisci un numero intero:")
	fmt.Scan(&n)
	if(n<1||n>10){
		fmt.Println("valore non valido")
	}else if(n==nfisso){
		fmt.Println("Hai indovinato!")
	}else{
		fmt.Println("Ciao")
	}
    
}
