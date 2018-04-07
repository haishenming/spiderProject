package main

import (
	"fmt"
	"spiderProject/tree"
)

type MyTreeNode struct {
	node *tree.TreeNode
}

func (myNode *MyTreeNode) postOrder() {
	if myNode == nil || myNode.node == nil {
		return
	}

	left := MyTreeNode{myNode.node.Left}
	left.postOrder()
	right := MyTreeNode{myNode.node.Right}
	right.postOrder()
	myNode.node.Print()
}

func main() {
	var root tree.TreeNode
	fmt.Println(root)

	root = tree.TreeNode{Value: 3}
	root.Left = &tree.TreeNode{}
	root.Right = &tree.TreeNode{5, nil, nil}
	root.Right.Left = new(tree.TreeNode)

	nodes := []tree.TreeNode{
		{Value: 3},
		{},
		{3, nil, nil},
	}
	fmt.Println(nodes)

	root.Print()

	root.Right.Left.SetValue(4)
	root.Right.Left.Print()

	root.Traverse()

	myRoot := MyTreeNode{&root}
	myRoot.postOrder()


}
