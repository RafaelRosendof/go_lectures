package main

import (
	"BStree/arvore"
	"fmt"
)

func main() {

	var raiz *arvore.Arv = nil

	raiz = arvore.Lendo_csv(raiz, "dados.csv")

	fmt.Println("Printando em ordem: ")

	arvore.Imprime_preOrdem(raiz)
}
