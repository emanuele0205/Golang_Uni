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
	var g,m,a int
	fmt.Println("giorno,mese,anno:")
	fmt.Scan(&g,&m,&a)
if a>=1582{
	if a%10!=0 && a%4==0{
		ris:=controllo_data_bis(g,m)
		if ris==true{
			fmt.Println("data valida")
		}else{
			fmt.Println("data non valida")
		}
	}else if a%10==0 && a%400==0{
		ris:=controllo_data_bis(g,m)
		if ris==true{
			fmt.Println("data valida")
		}else{
			fmt.Println("data non valida")
		}
	}else{
		ris:=controllo_data(g,m)
		if ris==true{
			fmt.Println("data valida")
		}else{
			fmt.Println("data non valida")
		}
	}
}else{
	fmt.Println("anno non valido")
}
}

func controllo_data (g int, m int) bool {
	var t bool
	if m==1 || m==3 || m==5 || m==7 || m==8 || m==10 || m==12{
		if g<=31{
			t=true
		}
		}else if m==4 || m==6 || m==9 || m==11{
			if g<=30{
				t=true
			}
		}else if m==2{
			if g<=28{
				t=true
			}
		}else{
			t=false
		}
		return t
		}


func controllo_data_bis (g int, m int) bool{
	var t bool
	if m==1 || m==3 || m==5 || m==7 || m==8 || m==10 || m==12{
		if g<=31{
			t=true
		}
		}else if m==4 || m==6 || m==9 || m==11{
			if g<=30{
				t=true
			}
		}else if m==2{
			if g<=29{
				t=true
			}
		}else{
			t=false
		}
		return t
		}

	
