package main

import (
	"fmt"
	"time"
)

func cookingChef(id int, c chan int) {
	fmt.Println("Chefe", id, " iniciou a cozinha")
	time.Sleep(time.Second * 5)
	c <- id
}

func main() {
	c := make(chan int)

	for i := 0; i < 5; i++ {
		go cookingChef(i, c) // inicia goroutine
	}

	for i := 0; i < 5; i++ {
		chefId := <-c // recebe o id pelo channel
		fmt.Println("Chefe", chefId, " terminou de cozinhar")
	} // terminou todas as goroutines
}
