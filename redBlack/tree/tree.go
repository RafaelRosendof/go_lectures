package tree

import "fmt"

/*
// Funções básicas
Tree* Arv_criaArv(); Feito
No* Arv_criaNo(No* pai, No node, Color color);
bool ArvoreVazia(Tree* arv);

// Funções específicas para Red-Black Tr
Tree* arv_insereRB(Tree* arv, No node); Feito
void corrigeInsercao(No* no); Feito
No* rotacaoEsquerda(Tree* arv, No* no); Feito
No* rotacaoDireita(Tree* arv, No* no); Feito
bool arv_removeRB(Tree* arv, int score);
void corrigeRemocao(No* no);

void ImprimeOrdem(No* raiz); Feito
void ImprimePreOrdem(No* raiz); Feito
No* Busca_no(No* raiz, int score);  feito
int Altura_Tree(No* raiz); feito

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

type Tree struct {
	raiz *Node
}

func Altura_Tree(raiz *Node) int {
	if raiz == nil {
		return -1
	}

	alt1 := Altura_Tree(raiz.esq) //checar isso
	alt2 := Altura_Tree(raiz.dir) // isso tbm

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

func ArvoreVazia(Tree *Tree) bool {
	if Tree.raiz == nil {
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

func Busca_no_raiz(raiz *Tree, score int) bool {
	if raiz == nil || raiz.raiz == nil {
		return false
	}

	if raiz.raiz.score == score {
		return true

	} else if raiz.raiz.score < score {
		return Busca_no_raiz(&Tree{raiz.raiz.dir}, score)
	} else {
		return Busca_no_raiz(&Tree{raiz.raiz.esq}, score)
	}
}

/*

func Busca_no(raiz *Tree, score int) *Node {
	if raiz == nil || raiz.raiz == nil {
		return nil
	}

	current := raiz.raiz

	for current != nil {
		if current.score == score {
			return current
		} else if current.score < score {
			current = current.dir
		} else {
			current = current.esq
		}
	}

	return nil
}
*/

func Arv_criaArv() *Tree {
	return &Tree{}
}

func PrintaNode(no *Node) {
	if no == nil {
		return
	}
	fmt.Println("Score: ", no.score, " Cor: ", no.cor)
}

func PrintaInOrdemNode(no *Node) {
	if no == nil {
		return

	}

	PrintaInOrdemNode(no.esq)
	PrintaNode(no)
	PrintaInOrdemNode(no.dir)
}

func PrintaInOrdem(raiz *Tree) {
	if raiz == nil || raiz.raiz == nil {
		return
	}
	PrintaInOrdemNode(raiz.raiz.esq)
	PrintaNode(raiz.raiz)
	PrintaInOrdemNode(raiz.raiz.dir)
}

func PrintaPreOrdem(raiz *Tree) {
	if raiz == nil || raiz.raiz == nil {
		return
	}
	PrintaNode(raiz.raiz)
	PrintaInOrdemNode(raiz.raiz.esq)
	PrintaInOrdemNode(raiz.raiz.dir)
}

func RotacaoEsq(arv *Tree, no *Node) *Node {
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

func RotacaoDir(arv *Tree, no *Node) *Node {

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

func Arv_insereRB(arv *Tree, no *Node) {
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
		} else if no.score > x.score {
			x = x.dir
		} else {
			return
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

func AjusteIRB(arv *Tree, no *Node) {
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

func transplant(arv *Tree, u, no *Node) {
	if u.pai == nil {
		arv.raiz = no
	} else if u == u.pai.esq {
		u.pai.esq = no
	} else {
		u.pai.dir = no
	}

	if no != nil {
		no.pai = u.pai
	}

	//no.pai = u.pai
}

func Minimo(no *Node) *Node {
	for no.esq != nil {
		no = no.esq
	}
	return no
}

/*
func corrigeRemocao(arv *Tree, no *Node) {
	for no != arv.raiz && (no == nil || no.cor == Black) {

		if no == nil {
			no = arv.raiz
		}

		if no.pai == nil {
			return
		}

		if no == no.pai.esq {
			w := no.pai.dir

			if w == nil {
				break
			}

			if w.cor == Red { //case 1 // error of segmentation fault here
				w.cor = Black
				no.pai.cor = Red // it was black here
				RotacaoEsq(arv, no.pai)
				w = no.pai.dir
			}

			//case 2
			if (w.esq == nil || w.esq.cor == Black) && (w.dir == nil || w.dir.cor == Black) {
				w.cor = Red
				no = no.pai
			} else {
				// case 3
				if w.dir == nil || w.dir.cor == Black {
					if w.esq != nil {
						w.esq.cor = Black
					}
					w.cor = Red
					RotacaoDir(arv, w)
					w = no.pai.dir
				}

				// case 4
				w.cor = no.pai.cor
				no.pai.cor = Black
				if w.dir != nil {
					w.dir.cor = Black
				}
				RotacaoEsq(arv, no.pai)
				no = arv.raiz
			}
		} else {
			w := no.pai.esq

			if w == nil {
				break
			}

			if w.cor == Red { //case 1

				w.cor = Black
				no.pai.cor = Red
				RotacaoDir(arv, no.pai)
				w = no.pai.esq
			}
			//case 2
			if (w.esq == nil || w.esq.cor == Black) && (w.dir == nil || w.dir.cor == Black) { //segfault here
				//if w.dir.cor == Black && w.esq.cor == Black {
				w.cor = Red
				no = no.pai
			} else {
				//case 3
				if w.esq == nil || w.esq.cor == Black {
					if w.dir != nil {
						w.dir.cor = Black
					}
					w.cor = Red
					RotacaoEsq(arv, w)
					w = no.pai.esq
				}
				// case 4
				w.cor = no.pai.cor
				no.pai.cor = Black
				if w.esq != nil {
					w.esq.cor = Black
				}
				RotacaoDir(arv, no.pai)

				no = arv.raiz
			}
		}
	}

	if no != nil {
		no.cor = Black
	}
}
*/

func Arv_removeRB(arv *Tree, score int) bool {
	z := Busca_no(arv.raiz, score)
	if z == nil {
		return false // Node not found
	}

	var x, parentOfX *Node
	y := z
	yOriginalColor := y.cor

	if z.esq == nil {
		x = z.dir
		transplant(arv, z, x)
		parentOfX = z.pai
	} else if z.dir == nil {
		x = z.esq
		transplant(arv, z, x)
		parentOfX = z.pai
	} else {
		y = Minimo(z.dir)
		yOriginalColor = y.cor
		x = y.dir
		if y.pai == z {
			parentOfX = y
		} else {
			parentOfX = y.pai
			transplant(arv, y, x)
			y.dir = z.dir
			y.dir.pai = y
		}
		transplant(arv, z, y)
		y.esq = z.esq
		y.esq.pai = y
		y.cor = z.cor
	}

	if yOriginalColor == Black {
		corrigeRemocao(arv, x, parentOfX)
	}

	return true
}

func corrigeRemocao(arv *Tree, no, parent *Node) {
	for no != arv.raiz && (no == nil || no.cor == Black) {
		if parent == nil {
			break // Should not happen if no is not root
		}

		// Case: 'no' is a left child (or the position of a nil left child)
		if no == parent.esq {
			w := parent.dir // Sibling
			if w == nil {
				break // Should not happen in a valid RB tree
			}

			if w.cor == Red { // Case 1
				w.cor = Black
				parent.cor = Red
				RotacaoEsq(arv, parent)
				w = parent.dir
			}

			// Case 2
			if (w.esq == nil || w.esq.cor == Black) && (w.dir == nil || w.dir.cor == Black) {
				w.cor = Red
				no = parent
				parent = no.pai
			} else {
				// Case 3
				if w.dir == nil || w.dir.cor == Black {
					if w.esq != nil {
						w.esq.cor = Black
					}
					w.cor = Red
					RotacaoDir(arv, w)
					w = parent.dir
				}
				// Case 4
				w.cor = parent.cor
				parent.cor = Black
				if w.dir != nil {
					w.dir.cor = Black
				}
				RotacaoEsq(arv, parent)
				no = arv.raiz // Problem solved, exit loop
			}
		} else { // Case: 'no' is a right child
			w := parent.esq // Sibling
			if w == nil {
				break
			}

			if w.cor == Red { // Case 1
				w.cor = Black
				parent.cor = Red
				RotacaoDir(arv, parent)
				w = parent.esq
			}

			// Case 2
			if (w.esq == nil || w.esq.cor == Black) && (w.dir == nil || w.dir.cor == Black) {
				w.cor = Red
				no = parent
				parent = no.pai
			} else {
				// Case 3
				if w.esq == nil || w.esq.cor == Black {
					if w.dir != nil {
						w.dir.cor = Black
					}
					w.cor = Red
					RotacaoEsq(arv, w)
					w = parent.esq
				}
				// Case 4
				w.cor = parent.cor
				parent.cor = Black
				if w.esq != nil {
					w.esq.cor = Black
				}
				RotacaoDir(arv, parent)
				no = arv.raiz
			}
		}
	}
	if no != nil {
		no.cor = Black
	}
}

// Testing function to check if the tree is balanced, and following the Red-Black properties
func IsValid(arv *Tree) bool {
	if arv == nil || arv.raiz == nil {
		return true
	}
	if arv.raiz.cor != Black {
		fmt.Println("Validation Error: Root is not black.")
		return false
	}
	if !isBST(arv.raiz, nil, nil) {
		return false
	}
	_, isValid := isValidRBTree(arv.raiz)
	if !isValid {
		return false
	}
	return true
}

// TODO
func isValidRBTree(node *Node) (blackHeight int, isValid bool) {

	if node == nil {
		return 1, true
	}
	if node.cor == Red {
		if (node.esq != nil && node.esq.cor == Red) || (node.dir != nil && node.dir.cor == Red) {
			fmt.Printf("Validation Error: Red node %d has a red child.\n", node.score)
			return 0, false
		}
	}
	leftBlackHeight, isLeftValid := isValidRBTree(node.esq)
	if !isLeftValid {
		return 0, false
	}
	rightBlackHeight, isRightValid := isValidRBTree(node.dir)
	if !isRightValid {
		return 0, false
	}
	if leftBlackHeight != rightBlackHeight {
		fmt.Printf("Validation Error: Black height mismatch at node %d. (Left: %d, Right: %d)\n", node.score, leftBlackHeight, rightBlackHeight)
		return 0, false
	}
	if node.cor == Black {
		leftBlackHeight++
	}
	return leftBlackHeight, true
}

func isBST(node *Node, minNode *Node, maxNode *Node) bool {
	if node == nil {
		return true
	}
	if maxNode != nil && node.score >= maxNode.score {
		fmt.Printf("BST Violation: Node %d is not less than max %d\n", node.score, maxNode.score)
		return false
	}
	if minNode != nil && node.score <= minNode.score {
		fmt.Printf("BST Violation: Node %d is not greater than min %d\n", node.score, minNode.score)
		return false
	}
	return isBST(node.esq, minNode, node) && isBST(node.dir, node, maxNode)
}
