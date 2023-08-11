package ejercicios

import (
	"fmt"
	"strconv"
)

var rst string
var num int

func Ejec(cad string) (int, string) {
	var err error
	num, err = strconv.Atoi(cad)

	fmt.Println(err)

	if num > 100 {
		rst = "Es mayor a 100"
	} else {
		rst = "Es menor a 100"
	}

	return num, rst
}
