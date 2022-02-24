package main

import (
	"fmt"
	"oo/mytree"
	"oo/tree"
)

func main() {
	var root tree.Node

	root = tree.Node{Value: 3}
	root.Left = &tree.Node{}
	root.Right = &tree.Node{Value: 5}
	root.Right.Left = new(tree.Node)
	root.Left.Right = tree.CreateNode(4)
	fmt.Println(root)
	fmt.Println(root.Right.Left)

	fmt.Println("-------------------")
	node := []tree.Node{
		{Value: 3},
		{},
		{6, nil, &root},
	}
	fmt.Println(node)

	fmt.Println("-------------------")
	root.Right.Left.SetValue(4)
	root.Right.Left.Loop()

	pRoot := &root
	pRoot.SetValue(100)
	pRoot.Loop()

	fmt.Println("-------------------")
	tree2 := mytree.Node{Data: pRoot}
	tree2.Loop()

	fmt.Println("-------------------")
	tree3 := mytree.MyEmbeddingNode{Node: pRoot}
	tree3.Loop()
	tree3.Node.Loop()
	//tree3 = pRoot
}
