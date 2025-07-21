package test

import (
	"redblack/tree"
)

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

func TestInsert100(t *testing.T) {

}

func TestInsert1_000_000(t *testing.T) {

}

func TestInsert1_000_000_000(t *testing.T) {

}

func TestRemove100(t *testing.T) {

}

func TestRemove1_000_000(t *testing.T) {

}

func ItsBalanced(t *testing.T) {

}

func TestInsertAndFindAndDelete(t *testing.T) {

}
