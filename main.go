package main

import "fmt"

 func main(){
	 var nome string
	 var senha string
	 fmt.Println(" digite seu nome de usuario:",nome )
	 fmt.Scan(&nome)
	 	if nome == "admin" {
			fmt.Println("digite sua senha:", senha)
			fmt.Scan(&senha) 
			if senha == "1234" {
				fmt.Println("acesso permitido")			
			}else {
			 fmt.Println("acesso negado")
		}
	}else {
		fmt.Println("acesso negado")} 
	}