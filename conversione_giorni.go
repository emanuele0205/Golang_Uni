/*Problema: Scrivere un programma Go conversione_giorni.go
che converta un numero specificato di giorni (fornito dall’utente)
in anni, seimane, giorni. Si ignorino gli anni bisestili.
Indicazione. Usare costanti per il numero di giorni in un anno e il
numero di giorni in una settimana.*/

package main
import "fmt"
func main(){

	const(
		 anno=365
		 sett=7
	)
	var giorni,tota,tots,totg int

	fmt.Println("inserisci il numero di giorni da convertire")
	fmt.Scan(&giorni)
	tota=giorni/anno
	tots=((giorni-(tota*anno))/sett)
	totg=giorni-(tota*anno)-(tots*sett)

	fmt.Println(giorni," giorni equivalgono a ",tota," anni, ",tots," settimane, ",totg," giorni")


}