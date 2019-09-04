/*Scrivere un programma lez5_giornodellanno.go che chiede 
all'utente due interi corrispondenti al giorno e mese di 
una data e calcola il numero del giorno dell'anno 
(cioe` il numero di giorni dall'inizio dell'anno).
*/
package main 
import "fmt"

func main(){
	var (
		giorno int 
		mese int 
		anno int
		sommagiorni int 
	)
		fmt.Println("inserire giorno e mese e anno")
		fmt.Scan(&giorno,&mese,&anno)
		if isDataValida(giorno,mese,anno)==true{
		sommagiorni=giorno
		for t:=1;t<mese;t++{
			sommagiorni+=numGiorniMese(t,anno)
		}
		
	}else{
		
		fmt.Println("data non valida ")
	}
if sommagiorni!=0{
fmt.Println("numero di giorni dall'inizio dell'anno ",sommagiorni)
}
		}


func numGiorniMese(mese int,anno int ) int{
	var numgiorni int

	controlloanno:=isBisestile(anno)
	if controlloanno==false{
	switch mese{
		case 1,3,5,7,8,10,12:
			numgiorni=31
		case 4,6,9,11:
			numgiorni=30
		case 2:
			numgiorni=28
	}
}else{
	switch mese{
	case 1,3,5,7,8,10,12:
		numgiorni=31
	case 4,6,9,11:
		numgiorni=30
	case 2:
		numgiorni=29
}
}
return numgiorni
}


func isDataValida(gg,mm,aa int ) bool{
	 var flag bool
	if gg>numGiorniMese(mm,aa)||gg<1||mm<1||mm>12{
		flag=false
	}else{
		flag=true
	}
	return flag
}

func isBisestile(anno int )bool{

	var flag bool
	if anno<4||anno==1581{
		flag=false
	}
	if anno>=4 && anno<=1580{
		if anno%4==0{
			flag=true
		}else{
			flag=false
		}
	}else{
		if anno%10!=0 && anno%4==0{
			flag=true
		}else if anno%10==0 && anno%400==0{
			flag=true
		}else{
			flag=false
		}
	}
	return flag	
}