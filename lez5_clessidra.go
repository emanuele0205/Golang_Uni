/*Scrivere un programma lez5_clessidra che mostri una clessidra 
(animata) che scandisce un countdown.*/

package main 
import(
	"os"
	"os/exec"
	"time"
	"fmt"
	"strings"
)


func main() {

// chiede dati a utente
	var sec int
	fmt.Println("Quanti secondi?")
	fmt.Scan(&sec)

// chiama la funzione disegna clessidra in un ciclo temporizzato,
// cancellando lo schermo tra un disegno e l'altro
	for i:= 0; i <= sec; i++{
		cancella()
		fmt.Println(i)
		clessidra(sec, i)
		attendi(1)
	}
}

 
// rigaClessidra disegna una singola riga
// accetta lunghezza, stato di piena o meno, carattere da usare, shift
// usando strings (senza usare il for)
func rigaClessidra(length int, full bool, char byte, shift int)  string{
	var row, offset, sand string
	offset = strings.Repeat(" ", shift)
	if full{
		sand = strings.Repeat(string(char), length)
	} else { // not full
		sand = string(char) + strings.Repeat(string(" "), length-2) + string(char)
	}
	row = offset + sand
	return row
}

// attendi attende
// time.Sleep(n_seconds * time.Second)
// (o time.Millisecond, time.Nanosecond....)
func attendi(n_seconds float64){
	time.Sleep(time.Duration(n_seconds) * time.Second)
}

// cancella fa il clear dello schermo
func cancella() {
	cmd := exec.Command("clear")
	cmd.Stdout=os.Stdout
	cmd.Run()
}

// clessidra disegna tutta la clessidra
// accetta altezza e secondi passati
func clessidra(height int, time int){
	width := height * 2 + 1
	full := false
	for i:=0; i < height; i++{
		if i >= time{
			full = true
		}
		fmt.Println( rigaClessidra(width -2*i, full, '*', i) )
	}
	full = false
	for i:=height-1; i >=0; i--{
		if time > i{
			full = true
		}
		fmt.Println( rigaClessidra(width -2*i, full, '*', i) )
	}
}