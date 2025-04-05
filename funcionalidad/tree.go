package funcionalidad

import (
	"os"
	"os/exec"
	"sistemagestoarchivos/helpers"
	"sistemagestoarchivos/operaciones"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func ControlarEnter(node *tview.TreeNode, app *tview.Application) {
	path := node.GetReference().(string)
	info, err := os.Stat(path)
	if err != nil {
		return
	}
	if info.IsDir() {
		if len(node.GetChildren()) == 0 {
			helpers.AddChildren(node, path)
		}
		node.SetExpanded(!node.IsExpanded())
	} else {
		app.Suspend(func() {
			cmd := exec.Command("vim", path)
			cmd.Stdin = os.Stdin
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			cmd.Run()
		})
	}
}

func CapturaOpcion(tree *tview.TreeView, app *tview.Application, event *tcell.EventKey, rootNode *tview.TreeNode) *tcell.EventKey {
	node := tree.GetCurrentNode()
	path := node.GetReference().(string)

	switch event.Rune() {
	case 'c':
		operaciones.Crear(path, tree, app, node)

	case 'd':
		operaciones.Eliminar(path, tree, app, rootNode)

	case 'r':
		operaciones.Renombrar(path, tree, app, rootNode)

	case 'm':
		operaciones.Renombrar(path, tree, app, rootNode)
	}
	return event
}
