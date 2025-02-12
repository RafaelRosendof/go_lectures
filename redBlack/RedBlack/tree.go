package redblack

import "fmt"

/*
// Funções básicas
tree* Arv_criaArv(); Feito
No* Arv_criaNo(No* pai, No node, Color color);
bool ArvoreVazia(tree* arv);

// Funções específicas para Red-Black Tr
tree* arv_insereRB(tree* arv, No node);
void corrigeInsercao(No* no);
No* rotacaoEsquerda(tree* arv, No* no);
No* rotacaoDireita(tree* arv, No* no);
bool arv_removeRB(tree* arv, int score);
void corrigeRemocao(No* no);

void ImprimeOrdem(No* raiz); Feito
void ImprimePreOrdem(No* raiz); Feito
No* Busca_no(No* raiz, int score);  feito
int Altura_tree(No* raiz); feito

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

func Altura_tree(raiz *Node) int {
	if raiz == nil {
		return 0
	}

	alt1 := Altura_tree(raiz.esq) //checar isso
	alt2 := Altura_tree(raiz.esq) // isso tbm

	if alt1 > alt2 {
		return alt1 + 1
	}
	return alt2 + 1
}

func Arv_criaNo(score int, pai *Node) *Node {
	return &Node{
		score: score,
		cor:   Red,
		esq:   nil,
		dir:   nil,
		pai:   pai,
	}
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

func RotacaoEsq(arv *tree, no *Node) *Node {
	filhoDir := no.dir
	no.dir = filhoDir.esq
	if filhoDir.esq != nil {
		filhoDir.esq.pai = no
	}
	filhoDir.pai = no.pai
	if no.pai == nil {
		arv.raiz = filhoDir
	} else if no == no.pai.esq {
		no.pai.esq = filhoDir
	} else {
		no.pai.dir = filhoDir
	}
	filhoDir.esq = no
	no.pai = filhoDir
	return filhoDir
}

func RotacaoDir(arv *tree, no *Node) *Node {

	filhoEsq := no.esq
	no.esq = filhoEsq.dir
	if filhoEsq.dir != nil {
		filhoEsq.dir.pai = no
	}

	filhoEsq.pai = no.pai
	if no.pai == nil {
		arv.raiz = filhoEsq
	} else if no == no.pai.dir {
		no.pai.dir = filhoEsq
	} else {
		no.pai.esq = filhoEsq
	}

	filhoEsq.dir = no
	no.pai = filhoEsq
	return filhoEsq

}

func Arv_insereRB(arv *tree, no *Node) {
	if arv.raiz == nil {
		no.cor = Black
		arv.raiz = no
		return
	}

	y := (*Node)(nil)
	x := arv.raiz

	for x != nil {
		y = x
		if no.score < x.score {
			x = x.esq
		} else {
			x = x.dir
		}
	}

	no.pai = y

	if no.score < y.score {
		y.esq = no
	} else {
		y.dir = no
	}

	no.esq = nil
	no.dir = nil
	no.cor = Red

	AjusteIRB(arv, no)

}

func AjusteIRB(arv *tree, no *Node) {
	for no != arv.raiz && no.pai != nil && no.pai.cor == Red {
		if no.pai == no.pai.pai.esq { //pai é filho esquerdo do avo
			y := no.pai.pai.dir //Tio é filho direito do avo

			if y != nil && y.cor == Red { //tio vermelho
				no.pai.cor = Black
				y.cor = Black
				no.pai.pai.cor = Red
				no = no.pai.pai
			} else {
				//no sendo fi Dir
				if no == no.pai.dir {
					no = no.pai
					RotacaoEsq(arv, no)
				}
				// FI esq
				no.pai.cor = Black
				no.pai.pai.cor = Red
				RotacaoDir(arv, no.pai.pai)
			}

		} else { // pai é filho direito do avo
			y := no.pai.pai.esq //tio filho esq do avo

			if y != nil && y.cor == Red {
				no.pai.cor = Black
				y.cor = Black
				no.pai.pai.cor = Red
				no = no.pai.pai //
			} else { //no é filho esq
				if no == no.pai.esq {
					no = no.pai
					RotacaoDir(arv, no)
				}

				//filho direito
				no.pai.cor = Black
				no.pai.pai.cor = Red
				RotacaoEsq(arv, no.pai.pai)
			}

		}
	}

	arv.raiz.cor = Black
	//acho que foi
}
