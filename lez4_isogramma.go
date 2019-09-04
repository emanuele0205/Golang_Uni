/*
Un isogramma e` una parola o una frase in cui nessuna
lettera e` ripetuta. Gli spazi, e i caratteri di underscore
e trattino e la punteggiatura possono invece apparire 
ripetuti.

Esempi di isogrammi:

lumberjacks
background
downstream
six-year-old

Scrivere un programma lez4_isogramma.go che legge una
parola e determina se e` un isogramma.*/

package main
import "fmt"
func main(){

	var(
		 s string 
		 alfabeto=[26]rune{'a','b', 'c', 'd' ,'e' ,'f' ,'g' ,'h' ,'i', 'j', 'k', 'l' ,'m' ,'n', 'o', 'p' ,'q', 'r', 's' ,'t' ,'u', 'v' ,'w' ,'x' ,'y' ,'z'} 
		 cont int 
		)
fmt.Println("inserire stringa")
fmt.Scan(&s)

	for i:=0;i<26;i++{
		for _,char:=range s{
		if char==alfabeto[i]{
			cont++
		}
	
}
if cont<=1{
		cont=0
		}
		if cont>=2{
			fmt.Println("la parola non e' un isogramma")
			break
		}
	}
	if cont<=1{
		fmt.Println("la parola e' un isogramma")
	}
}
