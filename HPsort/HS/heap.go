package HS

type Node struct{
	endOrigem	string
	endDestino	string
	size	int
	latencia	int
	prioridade	int
}

type Heap struct{
	tamanho		int
	capacidade		int
	pacotes		*[]Node
}

func CalculaPrioridade(n *Node)int{
	return -n.latencia + (n.size / 100)
}

//TODO o restante dos métodos
