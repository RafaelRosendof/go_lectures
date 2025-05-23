/*

bool remove_avl(tree *raiz , dado)

retorna concorrenteNode(no Node)



*/

package tree

import (
	"fmt"
)

type Node struct {
	Score int
	Year  int
	Figas int // integer bool 0 true 1 false
}

type Tree struct {
	node   Node
	left   *Tree
	rig    *Tree
	height int
}

func Tree_Empty(root *Tree) bool {
	return root == nil
}

func Tree_createNode(node Node) *Tree {
	newNode := &Tree{
		node:   node,
		left:   nil,
		rig:    nil,
		height: 0,
	}

	return newNode
}

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func balance_factor(root *Tree) int {
	if root == nil {
		return 0
	}

	return height_tree(root.left) - height_tree(root.rig)
}

func leftRout(root *Tree) *Tree {
	rootAux := root.rig

	root.rig = rootAux.left

	rootAux.left = root

	root.height = 1 + max(height_tree(root.left), height_tree(root.rig))
	rootAux.height = 1 + max(height_tree(rootAux.left), height_tree(rootAux.rig))
	return rootAux

}

func rightRout(root *Tree) *Tree {
	rootAux := root.left

	root.left = rootAux.rig

	rootAux.rig = root

	root.height = 1 + max(height_tree(root.left), height_tree(root.rig))
	rootAux.height = 1 + max(height_tree(rootAux.left), height_tree(rootAux.rig))

	return rootAux

}

func doubleLeft(root *Tree) *Tree {
	root.rig = rightRout(root.rig)
	return leftRout(root)
}

func doubleRight(root *Tree) *Tree {
	root.left = leftRout(root.left)
	return rightRout(root)
}

func Insert_avl(root *Tree, node Node) *Tree {

	if root == nil {
		return Tree_createNode(node)

	}

	if node.Score < root.node.Score {
		root.left = Insert_avl(root.left, node)

	} else if node.Score > root.node.Score {
		root.rig = Insert_avl(root.rig, node)
	} else {
		//fmt.Println("Node already in the tree")
		return root
	}

	root.height = 1 + max(height_tree(root.left), height_tree(root.rig))

	//update height and calculate the factor balance
	factor := balance_factor(root)

	//check for the balance

	if factor > 1 && node.Score < root.left.node.Score {
		return rightRout(root)
	}

	if factor < -1 && node.Score > root.rig.node.Score {
		return leftRout(root)
	}

	if factor > 1 && node.Score > root.left.node.Score {
		return doubleRight(root)
	}

	if factor < -1 && node.Score < root.rig.node.Score {
		return doubleLeft(root)
	}

	return root

}

func Remove_tree(root **Tree, Score int) {

	if root == nil {
		fmt.Println("Empty tree , nothing to remove")
		return
	}

	if *root == nil {
		fmt.Println("Node not found")
		return
	}

	if Score < (*root).node.Score {
		Remove_tree(&(*root).left, Score)
	} else if Score > (*root).node.Score {
		Remove_tree(&(*root).rig, Score)
	} else {

		//case with 0 childs

		if (*root).left == nil && (*root).rig == nil {
			*root = nil

		} else if (*root).left == nil { //case 1 child
			*root = (*root).rig

		} else if (*root).rig == nil {
			*root = (*root).left

		} else { //case with 2 childs

			rootAux := *root

			aux := (*root).left

			for aux.rig != nil {
				rootAux = aux
				aux = aux.rig
			}
			(*root).node = aux.node

			if rootAux == *root {
				rootAux.left = aux.rig
			} else {
				rootAux.rig = aux.left
			}
		}
	}

	if *root == nil {
		return
	}

	(*root).height = 1 + max(height_tree((*root).left), height_tree((*root).rig))
	factor := balance_factor(*root)

	if factor > 1 && balance_factor((*root).left) >= 0 {
		*root = rightRout(*root)
	}

	if factor > 1 && balance_factor((*root).left) < 0 {
		*root = doubleRight(*root)
	}

	if factor < -1 && balance_factor((*root).rig) <= 0 {
		*root = leftRout(*root)
	}

	if factor < -1 && balance_factor((*root).rig) > 0 {
		*root = doubleLeft(*root)
	}

	return
}

func Search_tree(root *Tree, Score int) Node {
	if root == nil {
		fmt.Println("Empty tree, nothing to serach ")
		return Node{}
	}

	if Score == root.node.Score {
		return root.node
	} else if root.node.Score > Score {
		return Search_tree(root.left, Score)
	} else if root.node.Score < Score {
		return Search_tree(root.rig, Score)
	}

	fmt.Println("The node with this value is not in the tree")
	return Node{}

}

func height_tree(root *Tree) int {
	if root == nil {
		return 0
	}

	h1 := height_tree(root.left)
	h2 := height_tree(root.rig)

	if h1 > h2 {
		return h1 + 1
	} else {
		return h2 + 1
	}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func IsBalanced(root *Tree) bool {
	if root == nil {
		return true
	}

	lh := height_tree(root.left)
	rh := height_tree(root.rig)

	if abs(lh-rh) > 1 {
		fmt.Printf("Node %d is unbalanced. Left height: %d, Right height: %d\n", root.node.Score, lh, rh)
		return false
	}

	return IsBalanced(root.left) && IsBalanced(root.rig)
}

func Print_order(root *Tree) {

	if root == nil {
		return
	}

	if root.left != nil {
		fmt.Println("My left son is : ", root.left.node.Score)
	} else {
		fmt.Println("I don't have a left son")
	}

	if root.rig != nil {
		fmt.Println("My right son is : ", root.rig.node.Score)
	} else {
		fmt.Println("I don't have a right son")
	}
	fmt.Println("Score: ", root.node.Score)
	fmt.Println("Year: ", root.node.Year)
	fmt.Println("It's Figas? : ", root.node.Figas)

	Print_order(root.left)
	Print_order(root.rig)

}

func Print_inorder(root *Tree) {

	Print_inorder(root.left)
	fmt.Println("Score: ", root.node.Score)
	fmt.Println("Year: ", root.node.Year)
	fmt.Println("It's Figas? : ", root.node.Figas)

	Print_inorder(root.rig)
}
