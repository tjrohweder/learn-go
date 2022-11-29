package main

import "fmt"

func fatorial(n uint) (uint, error) {
	switch {
	case n == 0:
		return 1, nil
	default:
		return fatorial(n - 1)
	}
}

func main() {
	resultado, _ := fatorial(5)
	fmt.Println(resultado)
}
