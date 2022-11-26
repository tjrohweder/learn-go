package main

import "fmt"

func main() {
	funcsESalarios := map[string]float64{
		"jose joao":      11325.45,
		"gabriela silva": 15456.78,
		"pedro junior":   1200.0,
	}

	funcsESalarios["Rafael Filho"] = 1350.0
	delete(funcsESalarios, "Inexistente")

	for nome, salario := range funcsESalarios {
		fmt.Println(nome, salario)
	}
}
