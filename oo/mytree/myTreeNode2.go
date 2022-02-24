package mytree

import (
	"fmt"
	"oo/tree"
)

type MyEmbeddingNode struct {
	*tree.Node
}

func (node *MyEmbeddingNode) Loop() {
	fmt.Println("Shadowed loop method.")
}
