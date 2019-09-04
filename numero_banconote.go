/*Problema. Scrivere un programma Go numero_banconote.go
che, dato un ammontare (valore intero) di denaro, lo suddivide
nel numero più piccolo possibile di banconote.
Annotazioni. Si supponga che i tagli di banconote disponibili
siano: 100, 50, 20, 10, 5, 2, 1.*/

package main
import "fmt"
func main(){
var amm,cento,cinquanta,venti,dieci,cinque,due,uno int
	fmt.Println("inserisci l'ammontare (numero intero)")
	fmt.Scan(&amm)
	
	cento=amm/100
	amm=amm-(cento*100)
	
	cinquanta=amm/50
	amm=amm-(cinquanta*50)

	venti=amm/20
	amm=amm-(venti*20)

	dieci=amm/10
	amm=amm-(dieci*10)

	cinque=amm/5
	amm=amm-(cinque*5)

	due=amm/2
	amm=amm-(due*2)

	uno=amm

	fmt.Println(cento,"banconota/e da 100")
	fmt.Println(cinquanta,"banconota/e da 50")
	fmt.Println(venti,"banconota/e da 20")
	fmt.Println(dieci,"banconota/e da 10")
	fmt.Println(cinque,"banconota/e da 5")
	fmt.Println(due,"banconota/e da 2")
	fmt.Println(uno,"banconota/e da 1")

	
	

}