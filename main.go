package main

import (
"fmt"

)

 func main(){
ages:=45
fmt.Println(ages<=50)
fmt.Println(ages >=50)
fmt.Println(ages ==50)
fmt.Println(ages !=50)

if ages <30{
	fmt.Println("menor que 30 anos")
}else if ages <40{
 fmt.Println("menor que 40 anos ")
}else {
	fmt.Println("não é menor que 40 anos")
}
names := []string{"Isadora", "Yasmim","Trunks", "Cebolinha", " Martin","Yuri"}

for index, value := range names {
	if index== 1{
		fmt.Println(" Continue apos a posição ", index, " e o valor ", value)
		continue 
	}

	if index >2 {
		fmt.Println("sair após", index)
		break
	}
	fmt.Println("Valor:",value)
}
 }