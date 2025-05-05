/*

bool remove_avl(tree *raiz , dado)

retorna concorrenteNode(no Node)


*/

package tree

import(
	"fmt"

)

type Node struct{
	score int
	name string
	year int
	figas bool
}

type Tree struct{
	dado Node
	left *Arv
	rig *Arv
	height int
}

func Tree_Empty(raiz *Arv) bool {
	return raiz == nil
}

func Tree_createNode(node Node) *Tree{
	newNode := &Tree{
		data: 	node
		left:	nil
		rig:	nil
		height:	0
	}

	return newNode
}

func max(a int , b int) int {
	if a > b{
		return a
	} else{
		return b
	}
}

func balance_factor(root *Tree) int {
	if root == nil{
		return 0
	}

	return height_tree(root.left) - height_tree(root.dir)
}

func leftRout(root *Tree) *Tree{
	rootAux := root.rig 

	root.rig = rootTemp.left

	rootAux.left = root

	root.height = 1 + max(height_tree(root.left) , height_tree(root.rig))
	rootAux.height = 1 + max(height_tree(rootTemp.left) , height_tree(rootTemp.rig))
	return rootAux

}

func rightRout(root *Tree) *Tree{
	rootAux := root.left 
	
	root.left = rootTemp.rig

	rootAux.rig = root 

	root.height = 1 + max(height_tree(root.left) , height_tree(root.rig))
	rootAux.height = 1 + max(height_tree(rootTemp.left) , height_tree(rootTemp.rig))

	return rootAux

}

func doubleLeft(root *Tree) *Tree{
	root.rig = rightRout(root.rig)
	return leftRout(root)
}

func doubleRight(root *Tree) *Tree{
	root.left = leftRout(root.left)
	return rightRout(root)
}


func Insert_avl(root *Tree , node Node) {

	if root == nil {
		return Tree_createNode(node)

	}

	if node.score < root.node.score {
		root.left = Insert_avl(root.left , node)

	}else if node.score > root.node.score{
		root.rig = Insert_avl(root.rig , node)
	}else{
		fmt.Println("Node already in the tree")
		return root
	}

	root.height = 1 + max(height_tree(root.left) , height_tree(root.rig))

	//update height and calculate the factor balance
	factor := balance_factor(root)

	//check for the balance

	if factor > 1 && node.score < root.left.node.score{
		return rightRout(root)
	}

	if factor < -1 && node.score > root.dir.node.score{
		return leftRout(root)
	}

	if factor > 1 && node.score > root.left.node.score{
		return doubleRight(root)
	}

	if factor < -1 && node.score < root.dir.node.score{
		return doubleLeft(root)
	}

	return root


}

func Remove_tree(root **Tree , node Node) bool{

	if root == nil{
		fmt.Println("Empty tree , nothing to remove")
		return false
	}

	if node.score < (*root).node.score{

	}
	else if node.score > (*root).node.score{

	}

	else{

		//case with 0 childs 


		//case with 1 child 



		//case with 2 childs 
	}

	root.height = 1 + max(height_tree((*root).left) , height_tree((*root).rig))
	factor := balance_factor(*root)

	if factor > 1 && balance_factor((*root).left) >= 0{
		*root = rightRout(*root)
	}

	if factor > 1 && balance_factor((*root).left) < 0{
		*root = doubleRight(*root)
	}

	if factor < -1 && balance_factor((*root).rig) <= 0{
		*root = leftRout(*root)
	}

	if factor < -1 && balance_factor((*root).rig) > 0{
		*root = doubleLeft(*root)
	}

	return true

}
func Search_tree(root *Tree , score int) Node{
	if raiz == nil{
		fmt.Println("Empty tree, nothing to serach ")
	}

	if root.score == score{
		return raiz
	} else if root.score < score{
		return Search_tree(root.esq , score)
	} else{
		return Search_tree(root.dir , score)
	}

	fmt.Println("The node with this value is not in the tree")
	return nil

}

func height_tree(root *Tree) int{
	if root == nil {
		return 0
	}

	h1 := height_tree(root.left)
	h2 := height_tree(root.rig)

	if h1 > h2{
		return h1 + 1
	} else{
		return h2 + 1
	}
}

func Print_order(root *Tree) {

	fmt.Println("Score: ", root.node.score)
	fmt.Println("Name: " , root.node.name)
	fmt.Println("Year: " , root.node.year)
	fmt.Println("It's figas? : " , root.node.figas)

	Print_order(root.left)
	Print_order(root.rig)

}

func Print_inorder(root *Tree){

	Print_inorder(root.left)
	fmt.Println("Score: ", root.node.score)
	fmt.Println("Name: " , root.node.name)
	fmt.Println("Year: " , root.node.year)
	fmt.Println("It's figas? : " , root.node.figas)

	Print_inorder(root.rig)
}
