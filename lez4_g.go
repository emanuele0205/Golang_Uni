/*Scrivere un programma lez4_g.go che stampa quanto segue
(fino a 25 caratteri)

G
GG
GGG 
GGGG 
GGGGG 
GGGGGG 
GGGGGGG 
...
*/

package main
import "fmt"
func main(){
	var n=25 
	var str3 string
	str1:='G'
	str2:=string(str1)
	for i:=0;i<n;i++{
		for j:=n;j>=n;j--{
			str3+=str2
			fmt.Print(str3)
	}
	fmt.Println()
}

}

