package main

import (
	"fmt"
)

func main() {
	var idade int

	// Pergunta a idade do usuário
	fmt.Print("Digite sua idade: ")
	fmt.Scan(&idade) // Lê a entrada do usuário

	// Verifica a faixa etária
	if idade < 18 {
		fmt.Println("Você é menor de idade.")
	} else if idade >= 18 && idade <= 60 {
		fmt.Println("Você é adulto.")
	} else {
		fmt.Println("Você é idoso.")
	}
}