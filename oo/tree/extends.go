package tree

import "fmt"

func (node *Node) Print() {
	if node == nil {
		return
	}
	fmt.Println(node.Value)
}

func (node *Node) Loop() {
	if node == nil {
		return
	}
	node.Print()
	node.Left.Loop()
	node.Right.Loop()
}

func (node *Node) SetValue(value int) {
	node.Value = value
}
