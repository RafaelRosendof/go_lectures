package redblack

import "fmt"

/*
// Funções básicas
tree* arv_criaVazia(); Feito
No* arv_criaNo(No* pai, No node, Color color);
bool arv_vazia(tree* arv);                Feito
// Funções específicas para Red-Black Tr
tree* arv_insereRB(tree* arv, No node);
void corrigeInsercao(No* no);
No* rotacaoEsquerda(tree* arv, No* no);
No* rotacaoDireita(tree* arv, No* no);
bool arv_removeRB(tree* arv, int score);
void corrigeRemocao(No* no);
void imprimeOrdem(No* raiz); Feito
void imprimePreOrdem(No* raiz); Feito
No* arv_busca(No* raiz, int score);  feito
int arv_altura(No* raiz);

ver uma lib de go para imprimir em formato de desenho
*/
type Color string

const (
	Red   Color = "Red"
	Black Color = "Black"
)

type Node struct {
	score int
	cor   Color
	esq   *Node
	dir   *Node
	pai   *Node
}

type tree struct {
	raiz *Node
}

func ArvoreVazia(tree *tree) bool {
	if tree.raiz == nil {
		return true
	}
	return false
}

func Max_num(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func Busca_no(raiz *Node, score int) *Node {
	if raiz == nil || raiz.score == score {
		return raiz
	}

	if raiz.score < score {
		return Busca_no(raiz.dir, score)
	}

	if raiz.score > score {
		return Busca_no(raiz.esq, score)
	}
	return nil
}

func Arv_criaArv() *tree {
	return &tree{}
}

func PrintaNode(no *Node) {
	if no == nil {
		return
	}
	fmt.Println("Score: ", no.score, " Cor: ", no.cor)
}

func PrintaInOrdem(raiz *Node) {
	if raiz == nil {
		return
	}
	PrintaInOrdem(raiz.esq)
	PrintaNode(raiz)
	PrintaInOrdem(raiz.dir)
}

func PrintaPreOrdem(raiz *Node) {
	if raiz == nil {
		return
	}
	PrintaNode(raiz)
	PrintaPreOrdem(raiz.esq)
	PrintaPreOrdem(raiz.dir)
}
