package test

//the github page for the test lib https://github.com/stretchr/testify

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"

	"AVL/tree"
)

func TestRoot(t *testing.T) {

	tests := []struct {
		name   string
		input  int
		output int
	}{
		{
			name:   "positive",
			input:  4,
			output: 16,
		},
		{
			name:   "negative",
			input:  -4,
			output: 16,
		},
		{
			name:   "positive",
			input:  5,
			output: 25,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			act := tree.Square(tt.input)
			assert.Equal(t, tt.output, act, "Square (%d) should be %d", tt.input, tt.output)
		})
	}
	t.Run("not Equal", func(t *testing.T) {
		act := tree.Square(25)
		assert.NotEqual(t, 624, act, "Not equal")
	})

}

func TestInsertAndFind(t *testing.T) {

	//get the data, insert in the tree and try to find the middle

	size := 200

	arx := make([]tree.Node, 0, size)

	for i := 0; i < size; i++ {
		score := rand.Intn(55) + 1
		year := rand.Intn(2025)
		figas := rand.Intn(10-5) + 10*4
		node := tree.Node{
			Score: score,
			Year:  year,
			Figas: figas,
		}
		arx = append(arx, node)

	}
	nodeRoot := tree.Node{
		Score: 15,
		Year:  2004,
		Figas: 1,
	}

	root := tree.Tree_createNode(nodeRoot)

	specificNode := tree.Node{
		Score: 3,
		Year:  2022,
		Figas: 42,
	}
	arx = append(arx, specificNode)

	for i := 1; i < size; i++ {
		root = tree.Insert_avl(root, arx[i])
	}

	key_score := 50
	wrong_score := 625

	t.Run("It find the score", func(t *testing.T) {
		act := tree.Search_tree(root, key_score)

		assert.Equal(t, key_score, act.Score, "The node was find in the tree ")
	})

	t.Run("Not find the score", func(t *testing.T) {
		act := tree.Search_tree(root, wrong_score)
		assert.NotEqual(t, act, wrong_score, "They should not be equal ")
	})
}

/*

Tests to do it latter

1 -> test for insert and find element in the tree ( insert + find )

2 -> test for insert and delete element in thee tree ( insert + delete )

3 ->

*/

/*

To make the correct test we need to put a file ex_test.go

and then use the go test

*/
