/*Scrivere un programma lez5_menu.go che mostra un menu 
con quattro scelte indicate con le lettere dalla a alla d. 
L'utente scrive la lettera corrispondente alla sua scelta 
e il programma mostra un messaggio che indica la scelta 
effettuata. 
Chiamare scelta la variabile in cui salvare l'opzione.

Esempio
-------

Proposta del giorno:
a. pizza
b. penne al pomodoro
c. cotoletta e patatine
d. crostata e caffe`

ordine?
	a
hai ordinato pizza*/

package main 
import "fmt"
func main(){

	var scelta string

	fmt.Print("a. pizza\nb. penne al pomodoro\nc. cotoletta e patatine\nd. crostata e caffe'\n\n")
	fmt.Println("ordine?")
	fmt.Scan(&scelta)
	fmt.Print("hai ordinato ")
	switch scelta{
		case "a":
			fmt.Print("pizza\n")
		case "b":
			fmt.Print("penne al pomodoro\n")
		case "c":
			fmt.Print("cotoletta e patatine\n")
		case "d":
			fmt.Print("crostata e caffe'\n")
		default:
			fmt.Println("scelta non valida")
			
		}

}