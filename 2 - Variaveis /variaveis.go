package main

import (
	"fmt"
)

// declarando variaveis
func main() {

	var nome string = "João" // maneira explicita de declarar uma variavel
	fmt.Println(nome)

	variavel2 := "Maria" // quando nao tem string, ele define pelo valor dela qual o tipo
	fmt.Println(variavel2)

	// declarando variaveis em bloco
	var (
		variavel3 string = "José"
		variavel4 string = "Ana"
	)
	fmt.Println(variavel3, variavel4)

	variavel5, variavel6 := "Pedro", "Paula" // declarando varias variaveis em uma linha
	fmt.Println(variavel5, variavel6)

	const constante string = "Sou uma constante" // constante nao pode ser alterada
	fmt.Println(constante)

}
