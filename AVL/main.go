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

}
