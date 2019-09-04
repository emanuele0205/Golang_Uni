/*Scrivere un programma Go lez2 esiti.go che associa voti in lettere a punteggi
secondo la seguente tabella:
90 - 100 A
80 - 89 B
70 - 79 C
60 - 69 D
0 - 59 F
Pi`u precisamente il programma riceve in ingresso un valore intero tra 0 e
100 e stampa su monitor il voto in lettere corrispondente.
Annotazioni Potendo dare per scontato che il valore voto fornito in ingresso
sar`a voto <= 100, se il primo if `e
if(90 <= voto)
come sar`a la codizione dell’else-if che lo segue? E necessario controllare che il `
voto sia minore di 90 oltre che maggiore o uguale a 80? Cio`e:
if(80 <= voto && voto < 90)
O basta verificare solo che il voto sia maggiore o uguale a 80? Cio`e:
if(80 <= voto)
Scegliere la soluzione che fa il minor numero possibile di controlli.*/

package main 
import "fmt"
func main(){
	var voto int 
	fmt.Println("voto?")
	fmt.Scan(&voto)
	if voto>=90 && voto<=100{
		fmt.Println("A")
	}else if voto>=80 && voto<=89{
		fmt.Println("B")
	}else if voto>=70 && voto<=79{
		fmt.Println("C")
	}else if voto>=60 && voto<=69{
		fmt.Println("D")
	}else if voto<=59{
		fmt.Println("F")
	}
}