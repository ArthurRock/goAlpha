package main

import (
	"fmt"
	"runtime"
	"strconv"

	"github.com/ArthurRock/goAlpha/ejercicios"
	"github.com/ArthurRock/goAlpha/files"
	"github.com/ArthurRock/goAlpha/iteraciones"
	"github.com/ArthurRock/goAlpha/variables"
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

	/*
		estado, texto := ejercicios.Ejec("679")
		fmt.Println(estado)
		fmt.Println(texto)

		teclado.IngresoNumero()
	*/

	// Imprimir los numeros

	variables.MuestroEnteros()

	variables.RestoVariables()

	estado, texto := variables.ConviertoText(23)
	fmt.Println(estado)
	fmt.Println(texto)

	if os := runtime.GOOS; os == "Linux." || os == "OS X." {
		fmt.Println("No es Windows, es ", os)
	} else {
		fmt.Println(runtime.GOARCH)
	}

	switch os := runtime.GOOS; os {
	case "linux":
		fmt.Println("Esto es Linux")
	case "darwin":
		fmt.Println("Esto es Darwin")
	default:
		fmt.Printf("Esto es %s \n", os)
	}

	intprint, strprint := ejercicios.Ejec("34")

	fmt.Println(strconv.FormatInt(int64(intprint), 10) + " " + strprint)

	iteraciones.Iterar(20)

	//fmt.Println(ejercicios.ValidarIngresoTeclado())

	//files.GrabaTabla()
	//files.SumaTabla()

	files.LeoArchivo2()

}
