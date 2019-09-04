/*Problema: Scrivere un programma Go altezza_rettangolo.go
che, date l’area e la base di un reangolo, determini l’altezza.
Annotazioni Verificare come si comporta il programma quando
l’area inserita non è un multiplo esao della base (nei due casi:
variabili dichiarate come int e come float64) e quando viene
immesso 0 come valore per la base.*/

package main 
import "fmt"
func main(){
	var area,base float64

	fmt.Println("inserire area e base")
	fmt.Scan(&area,&base)
	if base==0{
		fmt.Println("valore errato per la base")
		fmt.Println("inserire nuovamente valore della base")
		fmt.Scan(&base)
	}
	altezza:=area/base
	fmt.Println("l'altezza del rettangolo e':  ",altezza)
}