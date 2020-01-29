package main

import "fmt"

func main() {
	t := NewBinTree()
	t.Add(3)
	t.Add(2)
	t.Add(1)
	t.Add(-1)
	t.Add(5)

	// print values (sorted)
	t.Apply(func(val int) {
		fmt.Println(val)
	})

	// build sum
	ch := make(chan int)
	sum := 0
	go t.IterateValues(ch)
	for val := range ch {
		sum += val
	}
	fmt.Println("Sum =", sum)
}

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

func (node *BinTreeNode) Contains(elem int) bool {
	if elem == node.Elem {
		return true
	}
	if elem < node.Elem {
		if node.Left == nil {
			return false
		} else {
			return node.Left.Contains(elem)
		}
	}
	if node.Right == nil {
		return false
	} else {
		return node.Right.Contains(elem)
	}
}

type BinTree struct {
	Root *BinTreeNode
}

func NewBinTree() *BinTree {
	return &BinTree{Root: nil}
}

func (tree *BinTree) Add(elem int) {
	if tree.Root == nil {
		tree.Root = NewLeaf(elem)
	} else {
		tree.Root.Add(elem)
	}
}

func (tree *BinTree) Contains(elem int) bool {
	if tree.Root == nil {
		return false
	} else {
		return tree.Root.Contains(elem)
	}
}

func (node *BinTreeNode) IterateValues(ch chan<- int) {
	if node.Left != nil {
		node.Left.IterateValues(ch)
	}
	ch <- node.Elem
	if node.Right != nil {
		node.Right.IterateValues(ch)
	}
}

func (node *BinTreeNode) Apply(f func(val int)) {
	if node.Left != nil {
		node.Left.Apply(f)
	}
	f(node.Elem)
	if node.Right != nil {
		node.Right.Apply(f)
	}
}

func (tree *BinTree) IterateValues(ch chan<- int) {
	if tree.Root != nil {
		tree.Root.IterateValues(ch)
	}
	close(ch)
}

func (tree *BinTree) Apply(f func(val int)) {
	if tree.Root != nil {
		tree.Root.Apply(f)
	}
}
