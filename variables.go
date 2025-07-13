package main

import "fmt"

// declara uma variavel
var a int

// declara varias variaveis
var (
	b bool
	c float32
	d string
)

func main() {
	a = 123            // atribui um unico valor
	b, c = false, 1.23 // atribui varios valores
	d = "string"       // strings tem aspas duplas

	fmt.Println(a, b, c, d) // 123 false 1.23 string
}
