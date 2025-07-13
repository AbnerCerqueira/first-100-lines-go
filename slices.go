package main

import (
	"fmt"
	"reflect"
)

func main() {
	languages := [9]string{
		"C", "Lisp", "C++", "Java", "Python",
		"JavaScript", "Ruby", "Go", "Rust",
	}

	classicsLanguages := languages[0:3]  // a msm coisa q [:3]
	modernLanguages := make([]string, 4) // len(modern) = 4
	modernLanguages = languages[3:7]     // inclusivo:exclusivo
	newLanguages := languages[7:9]       // msm coisa q [7:]

	fmt.Printf("linguagens classicas %v\n", classicsLanguages)
	fmt.Printf("linguagens modernas %v\n", modernLanguages)
	fmt.Printf("linguagens novas %v\n", newLanguages)
	/*
		linguagens classicas [C Lisp C++]
		linguagens modernas [Java Python JavaScript Ruby]
		linguagens novas [Go Rust]
	*/

	// podemos pensar que slices são arrays dinamicos, arrays disfarçados
	allLanguages := languages[:]                     // copiando array
	fmt.Println(reflect.TypeOf(allLanguages).Kind()) // slice

	frameworks := []string{
		"React", "Vue", "Angular", "Svelte",
		"Laravel", "Django", "Flask", "Fiber",
	}

	jsFrameworks := frameworks[:4:4]              // tamanho atual 4 capacidade maxima 4
	jsFrameworks = append(jsFrameworks, "Meteor") // só é possível pois é um slice e não array

	frameworks = append(frameworks, jsFrameworks[4:]...)

	fmt.Printf("todas frameworks %v\n", frameworks)
	fmt.Printf("js frameworks %v\n", jsFrameworks)
	/*
		todas frameworks [React Vue Angular Svelte Laravel Django Flask Fiber Meteor]
		js frameworks [React Vue Angular Svelte Meteor]
	*/
}
