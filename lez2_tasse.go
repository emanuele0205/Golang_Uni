/*Scrivere un programma Go lez2 tasse.go che chiede reddito e stato civile (0
per non coniugato, 1 per coniugato) e calcola le tasse da pagare secondo la
seguente tabella:*/

package main
import "fmt"
func main(){

	var reddito,statciv,tot int
	fmt.Println("inserisci il tuo reddito")
	fmt.Scan(&reddito)
	fmt.Println("inserisci 0 per non coniugato o 1 per coniugato")
	fmt.Scan(&statciv)
	if(reddito>=0&&reddito<=32000&&statciv==0||reddito>=0&&reddito<=64000&&statciv==1){
		tot=(reddito*10)/100
	}else{
		tot=(reddito*25)/100
	}
fmt.Println("tasse: ",tot)
}