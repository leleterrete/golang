package main

import (
	"fmt"
)

func dadosPessoa(nome string, idade int) (int, string) {
    if idade > 18 {
        fmt.Println("Você é maior de idade")
    }else
    if idade <= 17{
        fmt.Println("Você é menor de idade") 
    }
    return idade, nome
  } 
  

  func main(){
    var idade int 
    var nome string

    fmt.Println("Digite sua idade:")
    fmt.Scan(&idade)
    fmt.Println("Sua idade é: ", idade)
    fmt.Println("Qual o seu nome?")
    fmt.Scan(&nome)
     dadosPessoa(nome, idade)
  }


  