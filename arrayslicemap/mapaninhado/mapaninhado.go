package main

import "fmt"

func main() {
	funcPorLetra := map[string]map[string]float64{
		"G": {
			"gabriela Silva": 15456.78,
			"guga pereira":   8456.78,
		},
		"J": {
			"jose joao": 11325.45,
		},
		"P": {
			"pedro junior": 1200.0,
		},

	delete(funcPorLetra, "P")

	for letra, funcs := range funcPorLetra {
		fmt.Println(letra, funcs)
	}
}
