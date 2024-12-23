package binaryTree

import (
	"fmt"
)

type Node struct {
	codigo     int
	titulo     string
	autor      string
	genero     string
	ano        int
	editora    string
	numPaginas int
}

type Arv struct {
	livro      Node
	num_filhos int
	esq        *Arv
	dir        *Arv
}

func arvoreVazia() *Arv {
	return nil
}

func arv_vazia(raiz *Arv) bool {
	return raiz == nil
}

func arv_criaNode(no Node) *Arv {

	novoNode := &Arv{
		livro:      no,
		esq:        nil,
		dir:        nil,
		num_filhos: 0,
	}
	return novoNode
}

func arvore_insere(raiz *Arv, no Node) *Arv {

	if raiz == nil {
		return arv_criaNode(no)
	}

	if no.codigo < raiz.livro.codigo {
		raiz.esq = arvore_insere(raiz.esq, no)
	} else if no.codigo > raiz.livro.codigo {
		raiz.dir = arvore_insere(raiz.dir, no)
	} else {
		fmt.Println("Livro já cadastrado: ", no.codigo)
	}

	return raiz
}

func arvore_busca(raiz *Arv, genero string) {

	if raiz == nil {
		fmt.Println("Arvore vazia não contém nada ")
		return
	}

	if raiz.livro.genero == genero {
		fmt.Println("Achou o livro EBA a posição é :", raiz.livro.codigo)
	}

	arvore_busca(raiz.esq, genero)
	arvore_busca(raiz.dir, genero)
}

func arvore_buscaCodigo(raiz *Arv, codigo int) *Arv {

	if raiz == nil || raiz.livro.codigo == codigo {
		return raiz
	}

	if raiz.livro.codigo < codigo {
		return arvore_buscaCodigo(raiz.dir, codigo)

	} else if raiz.livro.codigo > codigo {
		return arvore_buscaCodigo(raiz.esq, codigo)
	} else {
		fmt.Println("Não existe nenhum livro com esse código ")
		return nil
	}
}

// imprime pré ordem e pos ordem
func imprime_preOrdem(raiz *Arv) {
	if raiz == nil {
		return
	}

	fmt.Println() // to com preguiça de printar amanhã eu faço
}
