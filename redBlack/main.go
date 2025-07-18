package main

import (
	"fmt"
	"redBlack/tree"
)

func main() {

	fmt.Println("Initializing the RedBlack...")

	var root *tree.Tree = nil

	//tree.PrintaNode(root)
	tree.Write_csv("dataRed_black.csv")

	fmt.Println("Reading the CSV file ")

	root = tree.Read_csv(root, "dataRed_black.csv")

	fmt.Println("Printing in order: ")

	tree.PrintaInOrdem(root)

	//fmt.Println("Checking if the RedBlack is balanced: ")
	//balanced := tree.IsBalanced(root)
	//fmt.Println("Is the RedBlack balanced? ", balanced)

	// testing the remove for 1 element
	fmt.Println("Removing a node with score 512")

	//element := tree.Busca_no(root, 512)

	tree.Arv_removeRB(root, 512)

	fmt.Println("Remove worked :")

	fmt.Println("Choosing random numbers to remove -> ")
	remove_numbers := tree.Remove_nodes(500, "dataRed_black.csv")
	fmt.Println("Removing: ", remove_numbers[:10])

	for i := 0; i < len(remove_numbers); i++ {

		fmt.Println("Removing node with score: ", remove_numbers[i])

		if tree.Busca_no_raiz(root, remove_numbers[i]) == false {
			fmt.Println("Node with score", remove_numbers[i], "not found, skipping removal.")
			continue
		}

		tree.Arv_removeRB(root, remove_numbers[i])
	}

	fmt.Println("RedBlack after removals:")
	tree.PrintaInOrdem(root)

	fmt.Println("Checking if the RedBlack is balanced: ")
	//balanced := tree.IsBalanced(root)
	//fmt.Println("Is the RedBlack balanced? ", balanced)

	//Done,
}
