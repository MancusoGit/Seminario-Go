package ejercicios

import (
	"fmt"
)

func Imprimir250() {
	for i := 0; i <= 250; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
}
