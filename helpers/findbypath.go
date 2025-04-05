package helpers

import "github.com/rivo/tview"

func FindNodeByPath(node *tview.TreeNode, path string) *tview.TreeNode {
	if node.GetReference() == path {
		return node
	}
	for _, child := range node.GetChildren() {
		if found := FindNodeByPath(child, path); found != nil {
			return found
		}
	}
	return nil
}
