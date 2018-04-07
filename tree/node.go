package tree

import "fmt"

type TreeNode struct {
	Value       int
	Left, Right *TreeNode
}

func CreateNode(value int) *TreeNode {
	// 自定义构造函数
	return &TreeNode{Value: value}
}

func (node TreeNode) Print() {
	fmt.Print(node.Value, "\n")
}

func (node *TreeNode) SetValue(value int) {
	node.Value = value
}

func (node *TreeNode) Traverse() {
	if node == nil {
		return
	}
	node.Left.Traverse()
	node.Print()
	node.Right.Traverse()
}
