package main

import (
	"fmt"

	"github.com/ArthurRock/goAlpha/ejercicios"
)

func main() {
	/*
			estado, texto := variables.ConviertoText(34)
			fmt.Println(estado)
			fmt.Println(texto)

		if os := runtime.GOOS; os == "Linux." || os == "OS X." {
			fmt.Println("No es Windows, es ", os)
		} else {
			fmt.Println("Es Windows")
		}

		switch os := runtime.GOOS; os {
		case "linux":
			fmt.Println("Esto es Linux")
		case "darwin":
			fmt.Println("Esto es Darwin")
		default:
			fmt.Printf("%s \n", os)
		}
	*/
	estado, texto := ejercicios.Ejec("250")
	fmt.Println(estado)
	fmt.Println(texto)

}
