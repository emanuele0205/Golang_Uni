/*1.0.8 anno bisestile
Scrivere un programma Go lez2 bisestile.go che legge un intero n corrispondente all’anno di una data, e determina se l’anno `e bisestile o no.
Annotazioni Il calendario gregoriano si applica dal 1582, anno della sua
introduzione. Bench´e, in via teorica, sia possibile estenderlo anche agli anni
precedenti, per questi, di norma, si usa il calendario giuliano.
Sono bisestili tutti gli anni divisibili per 4, compresi quelli secolari, dal 4 al
1580. Per gli anni precedenti non si applicano gli anni bisestili. Per gli anni dal
1582 (calendario gregoriano) sono bisestili:
• gli anni non secolari il cui numero `e divisibile per 4;
• gli anni secolari il cui numero `e divisibile per 400.
Per esempio, il 1896 e il 1996 sono stati entrambi bisestili (non secolari
divisibili per 4), il 1800 e il 1900 non lo sono stati (secolari non divisibili per
400), mentre il 1600 e il 2000 lo sono stati (secolari divisibili per 400).*/

package main 
import "fmt"
func main(){
	var n int
	fmt.Println("anno:")
	fmt.Scan(&n)
	if n<4||n==1581{
		fmt.Println("impossibile calcolare se l'anno e' bisestile")
		fmt.Scan(&n)
	}
	if n>=4 && n<=1580{
		if n%4==0{
			fmt.Println("bisestile")
		}else{
			fmt.Println("non bisestile")
		}
	}else{
		if n%10!=0 && n%4==0{
			fmt.Println("bisestile")
		}else if n%10==0 && n%400==0{
			fmt.Println("bisestile")
		}else{
			fmt.Println("non bisestile")
		}
	}
		
}