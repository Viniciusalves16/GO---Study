package main

import (
	"fmt"
	"modulo/auxiliar"

	"github.com/badoux/checkmail" // importando o pacote auxiliar do pacote 1
)

// função main do pacote 1
func main() {
	fmt.Println("escrevendo do arquivo main do pacote 1")
	auxiliar.Escrever()
	erro := checkmail.ValidateFormat("em222")
	fmt.Println(erro)
}
