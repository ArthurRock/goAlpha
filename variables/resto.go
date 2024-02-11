package variables

import (
	"fmt"
	"strconv"
	"time"
)

var Nombre string
var Estado bool
var Sueldo float32
var Fecha time.Time

func RestoVariables() {

	Nombre = "Pedro"
	Estado = true
	Sueldo = 25000.10
	Fecha = time.Now()
	fmt.Println(Nombre)
	fmt.Println(Estado)
	fmt.Println(Sueldo)
	fmt.Println(Fecha)
}

func ConviertoText(numero int) (bool, string) {

	texto := strconv.Itoa(numero)
	return false, texto

}
