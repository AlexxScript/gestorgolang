package helpers

import (
	"os"
	"path/filepath"

	"github.com/rivo/tview"
)

func AddChildren(node *tview.TreeNode, path string) {
	node.ClearChildren()           //limpia el arreglo
	files, err := os.ReadDir(path) //lee los elementos del directorio y los retorna
	if err != nil {
		node.AddChild(tview.NewTreeNode("Error leyendo dir"))
		return
	}

	//for que adiciona los elementos que se encuentran en el directorio a la vista o el arbol
	for _, file := range files {
		childPath := filepath.Join(path, file.Name())
		childNode := tview.NewTreeNode(file.Name()).SetReference(childPath)
		if file.IsDir() {
			childNode.SetColor(tview.Styles.SecondaryTextColor)
		}
		node.AddChild(childNode) //se adiciona el archivo o carpeta al nodo padre de manera grafica
	}
}
