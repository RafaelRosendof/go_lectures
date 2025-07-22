package test

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"redBlack/tree"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Remove_csv(csv_in string) {
	cmd := exec.Command("rm", csv_in)

	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(out.String())

	fmt.Println("Done, we deleted the csv")
}

/*
func TestRoot(t *testing.T) {

	tests := []struct {
		score int
	}{
		{score: 10},
		{score: 120},
		{score: 304},
		{score: 410},
		{score: 50},
		{score: 60},
		{score: 702},
	}

	root := tree.Arva_criaArv()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			act := tree.Arv_insereRB(root, &tree.Node{score: tt.score, cor: tree.Red})

			assert.Equal(t, tt.score, act.GetRoot().score, "Root score should be %d", tt.score)
			assert.Equal(t, tree.Black, act.GetRoot().cor, "Root color should be Black")

		})
	}

}

*/

func TestInsert100(t *testing.T) {

	lines_in := 110
	csv_in := "teste_100.csv"

	tree.Write_csv(csv_in, lines_in)

	fmt.Println("Creating csv for 100 test")
	root := tree.Read_csv(nil, csv_in)

	is_valid := tree.IsValid(root)

	//Remove_csv(csv_in)
	assert.True(t, is_valid, "Tree is balanced and tested ok")

}

func TestInsert1_000_000(t *testing.T) {

	lines_in := 1_000_010
	csv_in := "test_1_000_000.csv"

	tree.Write_csv(csv_in, lines_in)

	fmt.Println("Creating csv for 1_000_000")
	root := tree.Read_csv(nil, csv_in)

	is_valid := tree.IsValid(root)

	//Remove_csv(csv_in)

	assert.True(t, is_valid, "tree is balanced for 1_000_000")

}

func TestInsert1_000_000_000(t *testing.T) {

	lines_in := 100_000_010
	csv_in := "test_1_000_000_000.csv"

	tree.Write_csv(csv_in, lines_in)

	fmt.Println("Creating csv for 1_000_000_000")
	root := tree.Read_csv(nil, csv_in)

	is_valid := tree.IsValid(root)

	//Remove_csv(csv_in)

	assert.True(t, is_valid, "tree is balanced for 1_000_000_000")
}

func TestRemove100(t *testing.T) {

	root := tree.Read_csv(nil, "teste_100.csv")

	fmt.Println("Testando remoção de 100 itens")

	remove_numbers := tree.Remove_nodes(100, "teste_100.csv")
	for i := 0; i < len(remove_numbers); i++ {

		if tree.Busca_no_raiz(root, remove_numbers[i]) == false {
			continue
		}

		tree.Arv_removeRB(root, remove_numbers[i])
	}

	is_valid := tree.IsValid(root)

	Remove_csv("teste_100.csv")
	assert.True(t, is_valid, "Tree is balanced after 100 removes")
}

/*
func TestRemove1_000_000(t *testing.T) {

	root := tree.Read_csv(nil, "test_1_000_000.csv")

	fmt.Println("Testando remoção de 1_000_000 itens")

	remove_numbers := tree.Remove_nodes(500_000, "test_1_000_000.csv")
	for i := 0; i < len(remove_numbers); i++ {
		//fmt.Println("Attempting to remove:", remove_numbers[i])
		if tree.Busca_no_raiz(root, remove_numbers[i]) == false {
			continue
		}

		tree.Arv_removeRB(root, remove_numbers[i])
	}

	is_valid := tree.IsValid(root)

	Remove_csv("teste_1_000_000.csv")
	assert.True(t, is_valid, "Tree is balanced after 1_000_000 removes")

}
*/

func TestRemove1_000_000_000(t *testing.T) {

	root := tree.Read_csv(nil, "test_1_000_000_000.csv")

	fmt.Println("Testando remoção de 1_000_000_000 itens")

	remove_numbers := tree.Remove_nodes(100_000, "test_1_000_000_000.csv")
	for i := 0; i < len(remove_numbers); i++ {

		if tree.Busca_no_raiz(root, remove_numbers[i]) == false {
			continue
		}

		tree.Arv_removeRB(root, remove_numbers[i])
	}

	is_valid := tree.IsValid(root)

	Remove_csv("test_1_000_000_000.csv")
	assert.True(t, is_valid, "Tree is balanced after 1_000_000_000 removes")

}
