package ejercicios

import (
	"fmt"
	"strconv"
)

func ValidarIngresoTeclado() string {

	var inputText string
	var errdata = 0
	var texto string

	for errdata != 1 {

		fmt.Println("Ingresar valor numerico")
		fmt.Scanln(&inputText)

		num, err := strconv.Atoi(inputText)

		if err != nil {
			fmt.Printf("Hubo un error " + err.Error() + "\n")
			errdata = 0
		} else {
			fmt.Printf("Tabla de multiplicar del num " + strconv.Itoa(num) + "\n")
			errdata = 1
			for i := 1; i < 11; i++ {
				texto += fmt.Sprintf(strconv.Itoa(i) + " * " + strconv.Itoa(num) + " = " + strconv.Itoa(num*i) + "\n")
			}
			texto += "\n"
		}
	}

	return texto
}
