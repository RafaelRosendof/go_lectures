package main

//código para fazer uma lista simples e encadeada

import (
	"container/list" //acho que é a biblioteca para estruturas de dados em geral
	"fmt"
)

func main() {
	var lista list.List

	lista.PushBack(32)
	lista.PushBack(19)
	lista.PushBack(90)

	//percorrendo a lista

	for i := lista.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value.(int))
	}

	//agora lendo do terminal

	fmt.Println("lendo do terminal agora ")
	fmt.Println("Quantos dígitos vc quer adicionar: ")

	var quantidade int
	fmt.Scanln(&quantidade)

	fmt.Println("Você agora alocou uma lista de tamanho: ", quantidade)

	var novos_num []int

	for i := 0; i != quantidade; i++ {
		fmt.Println("Digite o proximo número indíce %d: ", i+1)
		var item int
		fmt.Scanln(&item)
		novos_num = append(novos_num, item)
	}

	//adicionando a lista
	for _, num := range novos_num {
		lista.PushBack(num)
	}

	//mostrando o resto
	fmt.Println("Lista nova ")
	for i := lista.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value.(int))
	}
}

//TODO
/*
Rescrever alguns métodos do container/list no caso popcack e front e pushback e pushfront, quantidade head e tail, next e mais alguns métodos
*/
