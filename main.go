package main

import (
	"fmt"

	"github.com/ArthurRock/goAlpha/variables"
)

func main() {
	estado, texto := variables.ConviertoText(34)
	fmt.Println(estado)
	fmt.Println(texto)
}
