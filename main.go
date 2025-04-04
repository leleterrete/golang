package main

import (
"fmt"
"strings"
"sort"
)

 func main(){

   greeting := "Hello my friends!"
   fmt.Println(strings.Contains(greeting,"dog"))
fmt.Println(strings.ReplaceAll(greeting,"my ","mine "))
fmt.Println(strings.ToUpper(greeting))
fmt.Println(strings.Index(greeting,"my"))
fmt.Println(strings.Split(greeting," "))
ages := []int{50,80,10}
sort.Ints(ages)
fmt.Println(ages)
index := sort.SearchInts(ages,50)
fmt.Println(index)
names :=[]string{"Caroline", "Maicon", "Diego"}
sort.Strings(names)
fmt.Println(names)
fmt.Println(sort.SearchStrings(names,"Caroline"))
}

