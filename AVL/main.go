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

	root = nil
}
