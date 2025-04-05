package helpers

import (
	"os"
	"path/filepath"

	"github.com/rivo/tview"
)

func AddChildren(node *tview.TreeNode, path string) {
	node.ClearChildren()
	files, err := os.ReadDir(path) //lee los elementos del archivo y los retorna
	if err != nil {
		node.AddChild(tview.NewTreeNode("Error leyendo dir"))
		return
	}

	for _, file := range files {
		childPath := filepath.Join(path, file.Name())
		childNode := tview.NewTreeNode(file.Name()).SetReference(childPath)
		if file.IsDir() {
			childNode.SetColor(tview.Styles.SecondaryTextColor)
		}
		node.AddChild(childNode)
	}
}
