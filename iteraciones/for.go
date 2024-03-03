package iteraciones

import (
	"fmt"
	"strconv"
)

func Iterar(n int) {

	for i := 0; i <= n; i++ {

		if i%2 == 0 {
			continue
		}
		fmt.Println(strconv.FormatInt(int64(i), 10) + " <<<--- ")
	}
}
