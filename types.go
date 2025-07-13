package main

import "fmt"

func main() {

	// User specified types
	const a rune = 12         // 32-bit integer
	const b float32 = 20.5    // 32-bit float
	var c complex128 = 1 + 4i // 128-bit complex number
	var d uint16 = 14         // 16-bit unsigned number

	// Default types
	n := 42             // int
	pi := 3.14          // float
	x, y := true, false // bool
	z := "linguagem massa"

	fmt.Printf("tipos especificados pelo usuário:\n %T %T %T %T\n", a, b, c, d)
	/*
		tipos especificados pelo usuário:
		 int32 float32 complex128 uint16
	*/

	fmt.Printf("tipos default:\n %T %T %T %T %T\n", n, pi, x, y, z)
	/*
		tipos default:
		 int float64 bool bool string
	*/
}
