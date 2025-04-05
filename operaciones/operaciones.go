package operaciones

import (
	"os"
	"path/filepath"
	"sistemagestoarchivos/helpers"
	"strings"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func Crear(path string, tree *tview.TreeView, app *tview.Application, node *tview.TreeNode) {
	input := tview.NewInputField()

	input.SetLabel("Nombre nuevo archivo/directorio (termina en / para dir): ").
		SetDoneFunc(func(key tcell.Key) {
			name := input.GetText()
			full := filepath.Join(path, name)
			if strings.HasSuffix(name, "/") {
				os.MkdirAll(full, 0755)
			} else {
				os.WriteFile(full, []byte(""), 0644)
			}
			helpers.AddChildren(node, path)
			app.SetRoot(tree, true).SetFocus(tree)
		})
	app.SetRoot(input, true).SetFocus(input)
}

func Eliminar(path string, tree *tview.TreeView, app *tview.Application, rootNode *tview.TreeNode) {
	os.RemoveAll(path)
	parentPath := filepath.Dir(path)
	parentNode := helpers.FindNodeByPath(rootNode, parentPath)
	if parentNode != nil {
		helpers.AddChildren(parentNode, parentPath)
	}
}

func Renombrar(path string, tree *tview.TreeView, app *tview.Application, rootNode *tview.TreeNode) {
	input := tview.NewInputField()
	input.SetLabel("Nuevo nombre: ").
		SetDoneFunc(func(key tcell.Key) {
			newName := input.GetText()
			newPath := filepath.Join(filepath.Dir(path), newName)
			os.Rename(path, newPath)
			parentPath := filepath.Dir(path)
			parentNode := helpers.FindNodeByPath(rootNode, parentPath)
			if parentNode != nil {
				helpers.AddChildren(parentNode, parentPath)
			}
			app.SetRoot(tree, true).SetFocus(tree)
		})
	app.SetRoot(input, true).SetFocus(input)
}

func Mover(path string, tree *tview.TreeView, app *tview.Application, rootNode *tview.TreeNode) {
	input := tview.NewInputField()
	input.SetLabel("Mover a (ruta completa): ").
		SetDoneFunc(func(key tcell.Key) {
			dest := input.GetText()
			newPath := filepath.Join(dest, filepath.Base(path))
			os.Rename(path, newPath)
			parentPath := filepath.Dir(path)
			parentNode := helpers.FindNodeByPath(rootNode, parentPath)
			if parentNode != nil {
				helpers.AddChildren(parentNode, parentPath)
			}
			app.SetRoot(tree, true).SetFocus(tree)
		})
	app.SetRoot(input, true).SetFocus(input)
}
