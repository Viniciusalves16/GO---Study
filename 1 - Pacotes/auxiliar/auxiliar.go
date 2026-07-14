package auxiliar

import (
	"fmt"
)

// função que escreve uma mensagem no console
func Escrever() { // manter nome maiusculo para que seja exportada e possa ser usada em outros pacotes
	fmt.Println("escrevendo do arquivo auxiliar do pacote 1")
}
