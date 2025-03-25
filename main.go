package main

import (
	"bufio"
	"fmt"
	"os"
	"seminario/ejercicios"
)

func main() {
	fmt.Println("Ejercicio 1 : ")
	ejercicios.Imprimir250()
	fmt.Println()
	fmt.Println("Ejercicio 2 : ")
	num := 78
	resultado := ejercicios.Evaluar(num)
	fmt.Println("El resultado de evaluar el número", num, "es:", resultado)
	fmt.Println()
	fmt.Println("Ejercicio 3 : ")
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Ingrese una frase: ")
	frase, _ := reader.ReadString('\n')
	fmt.Println()
	fmt.Println("frase original: ", frase)
	fmt.Println("frase modificada: ", ejercicios.Ocurrencias(frase, "hola", "puto"))
	fmt.Println()
}
