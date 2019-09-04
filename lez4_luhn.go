/*scrivere un programma lez4_luhn.go che legge una 
sequenza di cifre (ad es. il numero di una carta di 
credito) e ne controlla la validita` utilizzando
l'algoritmo di Luhn:
- sostituire una cifra no e una si' con il suo doppio, 
iniziando da destra
	- se il doppio e` un numero maggiore di 9, 
	  sottrarre 9 al doppio
- sommare tutte le cifre
- se la somma e` esattamente divisibile per 10, il 
numero e` valido*/

package main 
import(
	 "fmt"
	
) 
func main(){

	var(
		 s string
		 somma int 
	)
	
	fmt.Println("inserire codice carta di credito senza spazzi")
	fmt.Scan(&s)
	t:=len(s)	
	for i:=0;i<len(s);i++{
		flag:=int(s[t-1]*2)
		if flag>9{
			somma+=flag-9
		}else{
			somma+=flag
		}
		t--
	}
if somma%10==0{
			fmt.Println("numero di carta valida")
		}else{
			fmt.Println("numero di carta non valida")
		}
}
