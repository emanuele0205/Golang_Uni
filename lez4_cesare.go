/*Il metodo di cifratura di parole di Cesare e` basato 
sull'operazione di rotazione di una lettera. Ruotare 
una lettera c maiuscola di un valore rot significa 
traslare c di rot posti, ripartendo da 'A' quando si 
arriva alla 'Z'. Indicando con

c + rot = cRot

il fatto che ruotando la lettera c di rot posti si 
ottiene la lettera cRot, si ha:

A+1=B
A+3=D
D+2=F
Y+1=Z  
Y+2=A

Scrivere un programma lez4_cesare.go che legge una 
parola s e un numero r e produce la parola cifrata 
ruotando di r posti ciascuna lettera di s.
(Suggerimento: usare string(ch) per visualizzare 
il carattere)*/


package main
import "fmt"
func main(){

	var(
		s string
		r int
		alfabeto=[26]rune{'a','b', 'c', 'd' ,'e' ,'f' ,'g' ,'h' ,'i', 'j', 'k', 'l' ,'m' ,'n', 'o', 'p' ,'q', 'r', 's' ,'t' ,'u', 'v' ,'w' ,'x' ,'y' ,'z'} 
	)
 fmt.Println("inserire la parola e il numero di spostamenti")
 fmt.Scan(&s,&r)

 for _,char:=range(s){
	 for i:=0;i<len(alfabeto);i++{
		 if char==alfabeto[i]{
			 t:=i+r
			
			 if t>=26{
				 l:=t%26
				
			 fmt.Printf("%c",alfabeto[l])
		 }else{
			fmt.Printf("%c",alfabeto[t])
		 }
	 }
	 }
 }
	fmt.Println()
}



	