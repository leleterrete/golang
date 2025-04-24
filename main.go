package main

import (
	"fmt"
)
func analisarNotas(nota1, nota2 float64) (float64, string) {
  
media := ( nota1 + nota2)/2
var resultado string

    if media >= 6 {
        resultado = "Aprovado"
    } else {
        resultado = "Reprovado"
    }

    return media, resultado

}

func main() {
  var nota1 float64
  var nota2 float64
  fmt.Println("Digite sua primeira nota:")
  fmt.Scan(&nota1)
  fmt.Println("Digite sua segunda nota:")
  fmt.Scan(&nota2)
  media, resultado := analisarNotas(nota1 , nota2)
  fmt.Println("Média:", media)
  fmt.Println("Resultado:", resultado)
}
  

  