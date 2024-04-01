package files

import (
	//	"bufio"
	"bufio"
	"fmt"
	"io/ioutil"

	//	"ioutil"
	"os"

	"github.com/ArthurRock/goAlpha/ejercicios"
)

var fileName string = "./files/txt/tabla.txt"

func GrabaTabla() {
	var texto string = ejercicios.ValidarIngresoTeclado()
	archivo, err := os.Create(fileName)

	if err != nil {
		fmt.Println("Error al crear el archivo " + err.Error())
		return
	}

	fmt.Fprintln(archivo, texto)
	archivo.Close()
}

func SumaTabla() {
	var texto string = ejercicios.ValidarIngresoTeclado()
	if Append(fileName, texto) == false {
		fmt.Println("Error al concatenar contenido")
	}
}

func Append(filen string, texto string) bool {

	arch, err := os.OpenFile(fileName, os.O_WRONLY|os.O_APPEND, 0644)

	if err != nil {
		fmt.Println("Error durante el Append " + err.Error())
		return false
	}

	_, err = arch.WriteString(texto)

	if err != nil {
		fmt.Println("Error durante el WriteString " + err.Error())
		return false
	}

	arch.Close()

	return true
}

func LeoArchivo() {
	archivo, err := ioutil.ReadFile(fileName)

	if err != nil {
		fmt.Println("Error al leer archivo " + err.Error())
		return
	}
	fmt.Println(string(archivo))

}

func LeoArchivo2() {
	archivo, err := os.Open(fileName)

	if err != nil {
		fmt.Println("Error al leer archivo " + err.Error())
		return
	}

	scanner := bufio.NewScanner(archivo)
	for scanner.Scan() {
		registro := scanner.Text()
		fmt.Println("> " + registro)
	}

	archivo.Close()
	//fmt.Println(string(archivo))

}
