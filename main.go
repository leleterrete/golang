package main
 import "fmt"
 
 func main(){
	 var numero1 float32
	 var numero2 float32
	 fmt.Println("digite um numero :")
	 fmt.Scan(&numero1)
	 fmt.Println("digite outro numero:")
	 fmt.Scan(&numero2)
	 fmt.Println("a soma dos dois numeros é:", numero1 + numero2)
	 fmt.Println("a subtração dos numeros é :", numero1 - numero2)
	 fmt.Println("a multiplicação dos numeros é:",numero1 * numero2)
	 fmt.Println("a divisão dos numeros é:", numero1 / numero2)
 }

