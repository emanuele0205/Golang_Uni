/*Problema. Scrivere un programma Go
conversione_secondi.go che converta un numero dato di
secondi (fornito dall’utente) in giorni, ore, minuti, secondi.*/

package main
import "fmt"
func main(){
	var secin,ss,mm,hh,gg int

	fmt.Println("inserisci il numero di secondi da convertire")
	fmt.Scan(&secin)

	gg=secin/86400
	secin=secin-(gg*86400)

	hh=secin/3600
	secin=secin-(hh*3600)

	mm=secin/60
	secin=secin-(mm*60)

	ss=secin

	fmt.Println("g:h:m:s - ",gg,":",hh,":",mm,":",ss)

}