package main

import "fmt"

func main() {
	var address *int  // declarar ponteiro
	number := 42      // int
	address = &number // guarda endereço de memória de `number`
	value := *address // desreferenciando

	fmt.Println(address)
	fmt.Println(&number)
	fmt.Println(&value)
}
