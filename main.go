package main

import (
	"fmt"
)

func main() {
 estoque := map[string]int{
  "Coxinha": 10,
"Pão de queijo": 15,
"Refrigerante": 20,
  }
  fmt.Println("ESTOQUE:")
	for produto, quantidade := range estoque {
		fmt.Printf("%s: %d\n", produto, quantidade)
  }
  estoque["Coxinha"] -= 2
	estoque["Pão de queijo"] -= 1

  fmt.Println("ESTOQUE ATUALIZADO:")
	for produto, quantidade := range estoque {
		fmt.Printf("%s: %d\n", produto, quantidade)
  }
}