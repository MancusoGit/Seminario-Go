package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func cambiarMinusMayus(palabras, ocurrencia string) string {
	frase := strings.Fields(palabras) // crea un slice de las palabras
	var result strings.Builder
	ocurrencia = strings.TrimSpace(ocurrencia) // elimina espacios en blanco al principio y al final

	for i, palabra := range frase {
		if strings.EqualFold(palabra, ocurrencia) {
			permutarChars(&result, palabra, ocurrencia)
		} else {
			result.WriteString(palabra)
		}

		if i < len(frase)-1 {
			result.WriteString(" ") // se agregan espacios entre palabras
		}
	}
	return result.String()
}

func permutarChars(result *strings.Builder, palabraFrase, comparacionFrase string) {
	palabra := []rune(palabraFrase)
	comparacion := []rune(comparacionFrase)
	for i, char := range palabra {
		if palabra[i] == comparacion[i] {
			if unicode.ToLower(comparacion[i]) == palabra[i] {
				result.WriteRune(unicode.ToUpper(char)) // paso a mayuscula
			} else {
				result.WriteRune(unicode.ToLower(char)) // paso a minuscula
			}
		} else {
			if unicode.ToUpper(comparacion[i]) == palabra[i] {
				result.WriteRune(unicode.ToLower(char)) // paso a minuscula
			} else {
				result.WriteRune(unicode.ToUpper(char)) // paso a mayuscula
			}
		}
	}
}

func ProbarOcurrencias() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Ingrese una frase: ")
	frase, _ := reader.ReadString('\n')
	fmt.Println()
	fmt.Print("Ingrese la palabra a cambiar: ")
	ocurrencia, _ := reader.ReadString('\n')
	fmt.Println()
	fmt.Println("Frase original: ", frase)
	fmt.Println("Palabra a cambiar: ", ocurrencia)
	fmt.Println("Frase modificada: ", cambiarMinusMayus(frase, ocurrencia))
}
