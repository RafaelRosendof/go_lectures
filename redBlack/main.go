package main

import (
	"fmt"
	"math/rand"
	"redBlack/tree"
)

func main() {

	//fmt.Println("Initializing the RedBlack...")
	//
	//var root *tree.Tree = nil
	//
	//tree.Write_csv("dataRed_black.csv", 9_000_000)
	//
	//fmt.Println("Reading the CSV file ")
	//
	//root = tree.Read_csv(root, "dataRed_black.csv")
	//
	//fmt.Println("Choosing random numbers to remove -> ")
	//remove_numbers := tree.Remove_nodes(5_000_000, "dataRed_black.csv")
	//fmt.Println("Removing: ", remove_numbers[:10])
	//
	//for i := 0; i < len(remove_numbers); i++ {
	//	if tree.Busca_no_raiz(root, remove_numbers[i]) == false {
	//		continue
	//	}
	//	tree.Arv_removeRB(root, remove_numbers[i])
	//}
	//
	//fmt.Println("Passing the remove nodes function ")
	//fmt.Println("RedBlack after removals:")
	//fmt.Println("Check if the tree is valid: ", tree.IsValid(root))

	lines_in := rand.Intn(10_000_000) + 1_000_000
	lines_out := rand.Intn(1_000_000) + 100_000
	csv_in := "dataRed_black.csv"

	Pipeline_rbtree(lines_in, lines_out, csv_in)

}

func Pipeline_rbtree(lines_in int, lines_out int, csv_in string) {

	fmt.Println("Creating the CSV file with ", lines_in, " lines")
	tree.Write_csv(csv_in, lines_in)

	fmt.Println("Reading the CSV file: ", csv_in)
	root := tree.Read_csv(nil, csv_in)

	fmt.Println("Removing ", lines_out, " nodes from the tree")
	remove_numbers := tree.Remove_nodes(lines_out, csv_in)

	for i := 0; i < len(remove_numbers); i++ {
		if tree.Busca_no_raiz(root, remove_numbers[i]) == false {
			continue
		}
		tree.Arv_removeRB(root, remove_numbers[i])
	}

	fmt.Println("Check if the tree is valid: ", tree.IsValid(root))
}
