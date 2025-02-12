package HS

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	codigo     int
	endOrigem  string
	endDestino string
	size       int
	latencia   int
	prioridade int
}

type Heap struct {
	tamanho    int
	capacidade int
	pacotes    []Node
}

func CalculaPrioridade(n *Node) int {
	return -n.latencia + (n.size / 100)
}

func Heap_vazia(heap *Heap) bool {

	return heap == nil || heap.tamanho == 0

}

func Pertence_heap(heap *Heap, codigo int) bool {

	if Heap_vazia(heap) {
		return false
	}

	for i := 0; i < heap.tamanho; i++ {
		if heap.pacotes[i].codigo == codigo {
			fmt.Println("Achou a operação")
			return true
		}
	}

	fmt.Println("Nenhum elemento com esse código ")
	return false
}

func printaNode(no *Node) {
	fmt.Println("Codigo: ", no.codigo, " endereço de Origem: ", no.endOrigem, " Endereço de Destino: ", no.endDestino, " Tamanho: ", no.size, " Latência: ", no.latencia, " Prioridade: ", no.prioridade)
}

func Printa_heap(heap *Heap) {
	if Heap_vazia(heap) {
		fmt.Println("Heap está vazia, nada a printar ")
		return
	}

	for i := 0; i < heap.tamanho; i++ {

		printaNode(&heap.pacotes[i]) // verificar se é a referência mesmo
	}
}

func Printa_Unicode(heap *Heap, codigo int) {
	if heap == nil || heap.tamanho == 0 {
		fmt.Println("Heap vazia -> não pode printar nada ")
	}

	for i := 0; i < heap.tamanho; i++ {
		if heap.pacotes[i].codigo == codigo {
			printaNode(&heap.pacotes[i])
		}
	}

	fmt.Println("Nenhum processo encontrado com esse código ")
}

func heapfy_dow(heap *Heap, i int) {
	for {
		maior := i
		esq := i*2 + 1
		dir := i*2 + 2

		if esq < heap.tamanho && heap.pacotes[i].prioridade > heap.pacotes[maior].prioridade {
			maior = esq
		}

		if dir < heap.tamanho && heap.pacotes[i].prioridade > heap.pacotes[maior].prioridade {
			maior = dir
		}

		if maior == i {
			break
		}

		heap.pacotes[i], heap.pacotes[maior] = heap.pacotes[maior], heap.pacotes[i]
		i = maior
	}
}

func heapfy_up(heap *Heap, i int) {

	for i > 0 {
		pai := (i - 1) / 2

		if heap.pacotes[pai].prioridade >= heap.pacotes[i].prioridade {
			break
		}

		aux := heap.pacotes[pai]

		heap.pacotes[pai] = heap.pacotes[i]
		heap.pacotes[i] = aux

		i = pai
	}
}

func Remove_maior_prioridade(heap *Heap) bool {
	if heap == nil || heap.tamanho == 0 {
		fmt.Println("Heap vazia, nada a remover ")
		return false
	}

	//raiz := heap.pacotes[0]

	heap.pacotes[0] = heap.pacotes[heap.tamanho-1]
	heap.pacotes = heap.pacotes[:heap.tamanho-1]

	heap.tamanho--

	heapfy_dow(heap, 0)

	return true

}

func Inserir_heap(heap *Heap, no Node) {

	if heap.capacidade <= heap.tamanho {
		fmt.Println("Heap cheia, alocando mais espaço")
		heap.capacidade *= 2
	}

	no.prioridade = CalculaPrioridade(&no)

	heap.pacotes = append(heap.pacotes, no)
	heap.tamanho++

	heapfy_up(heap, heap.tamanho-1)

}

func Remover_heap(heap *Heap, codigo int) bool {
	if heap == nil || heap.tamanho == 0 {
		fmt.Println("Heap vazio, nada a remover")
		return false
	}

	for i := 0; i < heap.tamanho; i++ {
		if heap.pacotes[i].codigo == codigo {
			// Substitui o nó pelo último elemento
			heap.pacotes[i] = heap.pacotes[heap.tamanho-1]
			heap.pacotes = heap.pacotes[:heap.tamanho-1]
			heap.tamanho--

			// Reorganiza o heap
			heapfy_dow(heap, i)
			return true
		}
	}

	fmt.Println("Nenhum elemento com esse código encontrado")
	return false
}

func Leitor_csv(heap *Heap, csv_file string) *Heap {

	arq, err := os.Open(csv_file)

	if err != nil {
		fmt.Println("Arquivo não encontrado: ")
		return nil
	}

	defer arq.Close()

	leitor := bufio.NewScanner(arq)

	first := true

	for leitor.Scan() {
		linha := leitor.Text()

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
		size, _ := strconv.Atoi(campos[3])
		latencia, _ := strconv.Atoi(campos[4])

		trecho := Node{
			codigo:     codigo,
			endOrigem:  campos[1],
			endDestino: campos[2],
			size:       size,
			latencia:   latencia,
		}

		//prioridade := CalculaPrioridade(&trecho)

		Inserir_heap(heap, trecho)

	}
	if err := leitor.Err(); err != nil {
		fmt.Println("Erro ao ler o arquivo: ", err)
	}
	return heap
}
