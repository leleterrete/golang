package main

import (
	"fmt"
)
func pegarNome()(string,string){
return "Barry", "Allen"

}
  func main(){
    nome, sobrenome := pegarNome()
    fmt.Println(nome)
    fmt.Println(sobrenome)
  }


  