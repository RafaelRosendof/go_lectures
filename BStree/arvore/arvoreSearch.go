package arvore

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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

func ArvoreVazia() *Arv {
	return nil
}

func Arv_vazia(raiz *Arv) bool {
	return raiz == nil
}

func Arv_criaNode(no Node) *Arv {

	novoNode := &Arv{
		livro:      no,
		esq:        nil,
		dir:        nil,
		num_filhos: 0,
	}
	return novoNode
}

func Arvore_insere(raiz *Arv, no Node) *Arv {

	if raiz == nil {
		return Arv_criaNode(no)
	}

	if no.codigo < raiz.livro.codigo {
		raiz.esq = Arvore_insere(raiz.esq, no)
	} else if no.codigo > raiz.livro.codigo {
		raiz.dir = Arvore_insere(raiz.dir, no)
	} else {
		fmt.Println("Livro já cadastrado: ", no.codigo)
	}

	return raiz
}

func Arvore_busca(raiz *Arv, genero string) {

	if raiz == nil {
		fmt.Println("Arvore vazia não contém nada ")
		return
	}

	if raiz.livro.genero == genero {
		fmt.Println("Achou o livro EBA a posição é :", raiz.livro.codigo)
	}

	Arvore_busca(raiz.esq, genero)
	Arvore_busca(raiz.dir, genero)
}

func Arvore_buscaCodigo(raiz *Arv, codigo int) *Arv {

	if raiz == nil || raiz.livro.codigo == codigo {
		return raiz
	}

	if raiz.livro.codigo < codigo {
		return Arvore_buscaCodigo(raiz.dir, codigo)

	} else if raiz.livro.codigo > codigo {
		return Arvore_buscaCodigo(raiz.esq, codigo)
	} else {
		fmt.Println("Não existe nenhum livro com esse código ")
		return nil
	}
}

// imprime pré ordem e pos ordem
func Imprime_preOrdem(raiz *Arv) {
	if raiz == nil {
		return
	}

	fmt.Println("Código: ", raiz.livro.codigo)
	fmt.Println("Título: ", raiz.livro.titulo)
	fmt.Println("Autor: ", raiz.livro.autor)
	fmt.Println("Genêro: ", raiz.livro.genero)
	fmt.Println("Ano: ", raiz.livro.ano)
	fmt.Println("Editora: ", raiz.livro.editora)
	fmt.Println("Número de páginas: ", raiz.livro.numPaginas)

	Imprime_preOrdem(raiz.esq)
	Imprime_preOrdem(raiz.dir)
}

func Imprime_posOrdem(raiz *Arv) {

	Imprime_posOrdem(raiz.esq)
	Imprime_posOrdem(raiz.dir)

	fmt.Println("Código: ", raiz.livro.codigo)
	fmt.Println("Título: ", raiz.livro.titulo)
	fmt.Println("Autor: ", raiz.livro.autor)
	fmt.Println("Genêro: ", raiz.livro.genero)
	fmt.Println("Ano: ", raiz.livro.ano)
	fmt.Println("Editora: ", raiz.livro.editora)
	fmt.Println("Número de páginas: ", raiz.livro.numPaginas)

}

//TODO -> removeNode , arvore_altura , carregar o csv

func Arvore_altura(raiz *Arv) int {
	if raiz == nil {
		return 0
	}

	h1 := Arvore_altura(raiz.esq)
	h2 := Arvore_altura(raiz.dir)

	if h1 > h2 {
		return h1 + 1
	} else {
		return h2 + 1
	}
}

func Remove_node(raiz **Arv, codigo int) bool {

	if raiz == nil {
		fmt.Println("Arvore nula, não remove nada ")
		return false
	}

	if codigo < (*raiz).livro.codigo {
		return Remove_node(&(*raiz).esq, codigo)
	} else if codigo > (*raiz).livro.codigo {
		return Remove_node(&(*raiz).dir, codigo)
	} else {

		//caso com 0 fihlos
		if (*raiz).esq == nil && (*raiz).dir == nil {

			*raiz = nil

			return true
		}

		//caso com 1 filho
		if (*raiz).esq == nil || (*raiz).dir == nil {

			aux := (*raiz).esq

			if aux == nil {
				aux = (*raiz).dir
			}

			*raiz = aux
			return true
		}

		//caos com dois filhos

		aux := (*raiz).dir
		if aux.esq != nil {
			aux = aux.esq
		}
		(*raiz).livro = aux.livro
		return Remove_node(&(*raiz).dir, aux.livro.codigo)
	}
}

func Lendo_csv(raiz *Arv, csv_file string) *Arv {
	arq, err := os.Open(csv_file)

	if err != nil {
		fmt.Println("Arquivo não encontrado: ", err)
		return nil
	}

	defer arq.Close()

	scanner := bufio.NewScanner(arq)
	first := true

	for scanner.Scan() {
		linha := scanner.Text()

		if first {
			first = false
			continue
		}

		campos := strings.Split(linha, ",")
		if len(campos) < 7 {
			fmt.Println("Linha inválida: ", linha)
			continue
		}

		codigo, _ := strconv.Atoi(campos[0])
		ano, _ := strconv.Atoi(campos[4])
		num_paginas, _ := strconv.Atoi(campos[6])

		livro := Node{
			codigo:     codigo,
			titulo:     campos[1],
			autor:      campos[2],
			genero:     campos[3],
			ano:        ano,
			editora:    campos[5],
			numPaginas: num_paginas,
		}

		raiz = Arvore_insere(raiz, livro)

	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Erro ao ler o arquivo: ", err)
	}

	return raiz

}
