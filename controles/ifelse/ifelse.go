package main

import "fmt"

func imprimirresultado(nota float64) {
	if nota >= 7 {
		fmt.Println("Aprovado com nota", nota)
	} else {
		fmt.Println("Reprovado com nota", nota)
	}
}

func main() {
	imprimirresultado(7.3)
	imprimirresultado(5.1)
}
