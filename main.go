package main

import (
"fmt"

)

 func main(){
var ages[5]int 
fmt.Println("digite o primeiro numero: ")
fmt.Scan(&ages[0])
fmt.Println("digite o segundo numero: ")
fmt.Scan(&ages[1])
fmt.Println("digite o terceiro numero: ")
fmt.Scan(&ages[2])
fmt.Println("digite o quarto numero: ")
fmt.Scan(&ages[3])
fmt.Println("digite o quinto numero: ")
fmt.Scan(&ages[4])
fmt.Println(ages)
 var soma int
 soma = ages[0]+ages[1]+ages[2]+ages[3]+ages[4]
fmt.Println("a soma é:",soma)
 }