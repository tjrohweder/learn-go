package main

import (
	"fmt"
	"time"
)

// Channel (canal) - eh a forma de comunicacao entre goroutines
//canal eh um tipo

func doistresQuatroVezes(base int, c chan int) {
	time.Sleep(time.Second)
	c <- 2 * base // enviando dados para o canal

	time.Sleep(time.Second)
	c <- 3 * base

	time.Sleep(3 * time.Second)
	c <- 4 * base
}

func main() {
	c := make(chan int)
	go doistresQuatroVezes(2, c)

	a, b := <-c, <-c // recebendo os dados do canal
	fmt.Println(a, b)

	fmt.Println(<-c)
}
