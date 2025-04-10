package main

import (
"fmt"
)

func main(){
    var opcao string
    var saldo int
    var valor int
    var valor2 int
    var opcao2 string
    fmt.Println("digite seu saldo (um valor inteiro positivo): ")
    fmt.Scan(&saldo)
    for {
    fmt.Println("escolha 'sacar' ou 'depositar'")
    fmt.Scan(&opcao)
    if opcao == "sacar"{
        fmt.Println("quanto você quer sacar? ")
        fmt.Scan(&valor) }

     if valor <= saldo && opcao == "sacar" {
        fmt.Println("saldo suficiente")
        saldo = saldo - valor
        fmt.Println("seu saldo atual é ", saldo)
        fmt.Println("se você quiser continuar, digite 'continuar', se quiser sair digite 'sair' ")
        fmt.Scan(&opcao2)
    } else if valor > saldo && opcao == "sacar" {
        fmt.Println("saldo insuficiente, digite 'continuar' para tentar denovo ou 'sair' para terminar")
        fmt.Scan(&opcao2)
    }
    if opcao == "depositar"{
        fmt.Println("quanto você quer depositar?")
        fmt.Scan(&valor2)
        saldo = saldo + valor2
fmt.Println("seu saldo atual é ", saldo)
   fmt.Println("se você quiser continuar, digite 'continuar', se quiser sair digite 'sair' ")
   fmt.Scan(&opcao2)
    }
    if opcao2 == "continuar" {
        continue
    } else if opcao2 == "sair" {
        break
    }
}
}