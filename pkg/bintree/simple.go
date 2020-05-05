package bintree

import "fmt"

type SimpleNode struct {
	left  *SimpleNode
	right *SimpleNode
	value int
}

func NewSimpleTree(val int) *SimpleNode {
	return &SimpleNode{
		left:  nil,
		right: nil,
		value: val,
	}
}

func (n *SimpleNode) Insert(val int) {

	switch {

	// store value is greater than new one then go left
	case n.value > val:
		if n.left == nil {
			n.left = &SimpleNode{value: val}
		} else {
			n.left.Insert(val)
		}

	// store value is less than new one then go right
	case n.value < val:
		if n.right == nil {
			n.right = &SimpleNode{value: val}
		} else {
			n.right.Insert(val)
		}
	}

}

func (n *SimpleNode) Traverse() {

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
