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
	fmt.Println("frase modificada: ", ejercicios.Ocurrencias(frase, "hola", "auto"))
	fmt.Println()
	fmt.Println("Ejercicio 4 : ")
	fmt.Println("El resultado de la función MayorMenor es:", ejercicios.MayorMenor())
	fmt.Println("El resultado de la función MenorMayor es:", ejercicios.MenorMayor()) // no se puede dividir por 0 (si se ingresa un numero negativo da error)
	fmt.Println("El resultado de la función MayorMenorFlotante es:", ejercicios.MayorMenorFlotante())
}
