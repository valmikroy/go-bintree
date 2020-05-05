package bintree

import "fmt"

type Node struct {
	left  *Node
	right *Node
	value int
}

func NewTree(val int) *Node {
	return &Node{
		left:  nil,
		right: nil,
		value: val,
	}
}

func (n *Node) Insert(val int) {

	switch {

	// store value is greater than new one then go left
	case n.value > val:
		if n.left == nil {
			n.left = &Node{value: val}
		} else {
			n.left.Insert(val)
		}

	// store value is less than new one then go right
	case n.value < val:
		if n.right == nil {
			n.right = &Node{value: val}
		} else {
			n.right.Insert(val)
		}
	}

}

func (n *Node) Traverse() {

	// reach to rightmost value
	if n.right != nil {
		n.right.Traverse()
	}

	// start printing
	fmt.Println(n.value)

	// reach to leftmost value
	if n.left != nil {
		n.left.Traverse()
	}

}
