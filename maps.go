package main

import "fmt"

// a eficiencias de um map se da as custas de randomizar a ordem dos pares de chave-valor
func main() {
	firstReleases := map[string]int{
		"C": 1972, "C++": 1985, "Java": 1996,
		"Python": 1991, "JavaScript": 1996, "Go": 2012,
	}

	for k, v := range firstReleases {
		fmt.Printf("%s é de %d\n", k, v)
	}
}
