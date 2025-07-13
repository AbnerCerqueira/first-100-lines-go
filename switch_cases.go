package main

import "fmt"

func avg2(x []float64) (avg float64) {
	total := 0.0
	switch len(x) {
	case 0: // não é necessário usar `break` após o fim de `case`
		avg = 0.0
	default:
		for _, v := range x {
			total += v
		}
		avg = total / float64(len(x))
	}

	return
}

func main() {
	x := []float64{2.15} // slice do tipo float
	fmt.Println(avg2(x))
}
