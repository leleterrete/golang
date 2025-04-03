package main

import "fmt"

 func main(){
	var ages = [4]int{17,16,20,40}
	nomes := [4]string{"Mario", "Luigi", "DeadPool", "Superman"}
	fmt.Println(ages)
	fmt.Println(nomes)
    nomes[3] = " Clark Kent"
	fmt.Println(nomes)
}

