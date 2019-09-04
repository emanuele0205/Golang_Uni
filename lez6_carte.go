/*
- Definire un tipo struct carta con due campi di tipo string: 
valore ("A", "1", ..., "10", "J", "Q", "K") e seme 
("C", "Q", "F", "P").
- Definire una funzione estraiCarta() che genera un numero 
casuale tra 1 e 52 e restituisce la carta corrispondente 
(le prime 13 sono di cuori, dalla 14 alla 26 sono di quadri, 
ecc, quindi 11 -> fante di cuori; 14 -> asso di quadri)
- Definire una funzione main che chiami estraiCarta e stampi 
la carta estratta.
- Definire una funzione daiCarte(n int) che restituisca
una slice di n carte estratte con estraiCarta.
- Testarla nel main.*/

package main 
import "fmt"
func main(){

type carte struct{
	seme={"C","Q","F","P"}string
	valore={"A","1","2","3","4","5","6","7","8","9","10","J","Q","K"} string
}

}