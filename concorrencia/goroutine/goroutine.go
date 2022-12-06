package main

import (
	"fmt"
	"time"
)

func fale(pessoa, texto string, qtde int) {
	for i := 0; i < qtde; i++ {
		time.Sleep(time.Second)
		fmt.Printf("%s: %s (iteracao %d)\n", pessoa, texto, i+1)
	}
}

func main() {
	//fale("Maria", "pq vc nao fala comigo", 3)
	//fale("Maria", "so posso falar depois de vc", 1)

	go fale("Maria", "Ei", 500)
	go fale("Maria", "Opa", 500)

	time.Sleep(time.Second * 5)
	fmt.Println("Fim!")
}
