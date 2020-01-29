package main

type BinTreeNode struct {
	Elem        int
	Left, Right *BinTreeNode
}

func NewLeaf(elem int) *BinTreeNode {
	return &BinTreeNode{Elem: elem, Left: nil, Right: nil}
}

func (node *BinTreeNode) Add(elem int) {
	if elem == node.Elem {
		// do nothing, return
		return
	}
	if elem < node.Elem {
		if node.Left == nil {
			node.Left = NewLeaf(elem)
		} else {
			node.Left.Add(elem)
		}
	} else {
		if node.Right == nil {
			node.Right = NewLeaf(elem)
		} else {
			node.Right.Add(elem)
		}
	}
}

type BinTree struct {
	Root *BinTreeNode
}

func NewBinTree() *BinTree {
	return &BinTree{Root: nil}
}



func main() {

}
