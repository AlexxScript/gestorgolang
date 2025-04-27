package helpers

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/rivo/tview"
)

func AddChildren(node *tview.TreeNode, path string) {
	node.ClearChildren()

	files, err := os.ReadDir(path)
	if err != nil {
		node.AddChild(tview.NewTreeNode("Error leyendo dir"))
		return
	}

	for _, file := range files {
		childPath := filepath.Join(path, file.Name())
		info, err := file.Info()
		if err != nil {
			continue
		}
		perms := info.Mode().Perm()
		texto := fmt.Sprintf("%s (%s)", file.Name(), perms.String())
		childNode := tview.NewTreeNode(texto).SetReference(childPath)
		if file.IsDir() {
			childNode.SetColor(tview.Styles.SecondaryTextColor)
		}

		node.AddChild(childNode)
	}
}
