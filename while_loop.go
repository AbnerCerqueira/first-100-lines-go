package main

import "fmt"

// em go não existe a palavra `while`
func main() {
	count := 0

	for count < 10 {
		count += 1
		fmt.Println(count)
	}
}
