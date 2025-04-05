package funcionalidad

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/gdamore/tcell/v2"
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

func ControlarEnter(node *tview.TreeNode, app *tview.Application) {
	path := node.GetReference().(string)
	info, err := os.Stat(path)
	if err != nil {
		return
	}
	if info.IsDir() {
		if len(node.GetChildren()) == 0 {
			AddChildren(node, path)
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

func findNodeByPath(node *tview.TreeNode, path string) *tview.TreeNode {
	if node.GetReference() == path {
		return node
	}
	for _, child := range node.GetChildren() {
		if found := findNodeByPath(child, path); found != nil {
			return found
		}
	}
	return nil
}

func CapturaOpcion(tree *tview.TreeView, app *tview.Application, event *tcell.EventKey, rootNode *tview.TreeNode) *tcell.EventKey {
	node := tree.GetCurrentNode()
	path := node.GetReference().(string)

	switch event.Rune() {
	case 'c':
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
				AddChildren(node, path)
				app.SetRoot(tree, true).SetFocus(tree)
			})
		app.SetRoot(input, true).SetFocus(input)

	case 'd':
		os.RemoveAll(path)
		parentPath := filepath.Dir(path)
		parentNode := findNodeByPath(rootNode, parentPath)
		if parentNode != nil {
			AddChildren(parentNode, parentPath)
		}

	case 'r':
		input := tview.NewInputField()
		input.SetLabel("Nuevo nombre: ").
			SetDoneFunc(func(key tcell.Key) {
				newName := input.GetText()
				newPath := filepath.Join(filepath.Dir(path), newName)
				os.Rename(path, newPath)
				parentPath := filepath.Dir(path)
				parentNode := findNodeByPath(rootNode, parentPath)
				if parentNode != nil {
					AddChildren(parentNode, parentPath)
				}
				app.SetRoot(tree, true).SetFocus(tree)
			})
		app.SetRoot(input, true).SetFocus(input)

	case 'm':
		input := tview.NewInputField()
		input.SetLabel("Mover a (ruta completa): ").
			SetDoneFunc(func(key tcell.Key) {
				dest := input.GetText()
				newPath := filepath.Join(dest, filepath.Base(path))
				os.Rename(path, newPath)
				parentPath := filepath.Dir(path)
				parentNode := findNodeByPath(rootNode, parentPath)
				if parentNode != nil {
					AddChildren(parentNode, parentPath)
				}
				app.SetRoot(tree, true).SetFocus(tree)
			})
		app.SetRoot(input, true).SetFocus(input)
	}
	return event
}
