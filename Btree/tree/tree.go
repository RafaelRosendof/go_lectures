package tree

type BtreeNode struct {
	Num_keys int
	Keys     []int
	Leaf     bool
	Child    []*BtreeNode
}

type Btree struct {
	Root *BtreeNode
	T    int
}

/*
func (tree *BTree) Insert(k int) { ... }

func (n *BtreeNode) SplitChild(i int, t int) { ... }
*/
