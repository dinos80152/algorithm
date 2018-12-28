package binarytree

func InvertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	right := InvertTree(root.Right)
	left := InvertTree(root.Left)
	root.Left = right
	root.Right = left
	return root
}
