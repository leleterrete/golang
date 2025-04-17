package main

import (
	"fmt"
)

  func main(){
    alunoIdade := make(map[string]int)
    alunoIdade["Bruno"] = 15
    alunoIdade["Otavio"]= 16
    alunoIdade["Fabiano"] = 40
    alunoIdade["Isabela"] = 15
    fmt.Println("Idade do Bruno", alunoIdade["Bruno"])
    notasAlunos :=map[string]float64{
        "Bruno" : 9.7,
        "Otavio" :10,
        "Fabiano":8.7,
        "Isabela": 9.5,

    }
    for k, v:= range notasAlunos {
        fmt.Printf("%s tirou a nota % 1.f ", k ,v )
    }
  }


  