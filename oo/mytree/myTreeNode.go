package mytree

import "oo/tree"

type Node struct {
	Data *tree.Node
}

func (node Node) Loop() {
	if node.Data == nil {
		return
	}
	node.Data.Print()
	right := Node{Data: node.Data.Right}
	left := Node{Data: node.Data.Left}
	right.Loop()
	left.Loop()
}
