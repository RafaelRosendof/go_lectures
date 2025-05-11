package main

import (
	"AVL/tree"
	"fmt"
)

func main() {
	var root *tree.Tree = nil

	tree.Write_csv("data.csv")

	root = tree.Read_csv(root, "data.csv")

	fmt.Println("Printing in orderrer: ")

	tree.Print_order(root)

	fmt.Println("Checking if the tree is balanced: ")
	balanced := tree.IsBalanced(root)

	if balanced {
		fmt.Println("The tree is balanced")
	} else {
		fmt.Println("The tree is not balanced")
	}

	fmt.Println("Choosing random numbers to remove -> ")

	remove_numbers := tree.Remove_nodes(5_000, "data.csv")

	fmt.Println("Removing : ", remove_numbers)

	for i := 0; i < len(remove_numbers); i++ {
		tree.Remove_tree(&(root), remove_numbers[i])
	}

	fmt.Println("Printing in orderrer: ")

	tree.Print_order(root)

	fmt.Println("Checking if the tree is balanced: ")
	balanced2 := tree.IsBalanced(root)

	if balanced2 {
		fmt.Println("The tree is balanced")
	} else {
		fmt.Println("The tree is not balanced")
	}

	root = nil
}
