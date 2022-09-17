package main

import "fmt"

func main() {
	y := 1
	x := 2

	// apenas postfix
	x++ // x +=1 ou x= x+1
	fmt.Println(x)

	y-- //y -= 1 ou y = y -1
	fmt.Println(y)

	// fmt.Println(x++ == y--) Nao eh possivel
}
