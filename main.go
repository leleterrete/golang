package main

import (
	"fmt"
)

var saldo float64
var decisão1 float64
var newOPR float64

func decisão() {
    fmt.Println("O que você deseja fazer?\n1-depositar\n2-sacar\n3-versaldo")
    fmt.Scan(&decisão1)
    if decisão1 == 1 {
        addValor()
    }else if decisão1 == 2 {
        Saque()
    }else if decisão1 == 3 {
        VerSaldo()
    }else {
        fmt.Println("Ação inexistente ")
    }
}

func VerSaldo() {
    fmt.Println("Seu saldo é:", saldo)
}
func main() {
    saldo = 240
   fmt.Println("Seu saldo é:", saldo)
   decisão()
   for {
    fmt.Println("Você deseja realizar uma nova operação? 1-Sim 2-Não")
    fmt.Scan(&newOPR)
    if newOPR == 1 {
        decisão()
    }else if newOPR==2 { 
        break
    }else { 
        fmt.Println("Ação invalida")
    }
   } 
}
func addValor(){
    var dep float64
fmt.Println("Qual valor você quer depositar?")
fmt.Scan(&dep) 
if dep <= 0 {
    fmt.Println("Impossivel depositar valores menores ou iguais a 0")
}else {
    saldo = saldo + dep 
    fmt.Println("Seu novo saldo é:", saldo )
}

}

  func Saque(){
    var saque float64
    fmt.Println("Qual valor você quer sacar?")
fmt.Scan(&saque)
if saque <= 0 || saque > saldo {
    fmt.Println("É impossivel sacar")
}else {
    saldo = saldo - saque 
    fmt.Println("Seu novo saldo é", saldo )
}
  
  }


  