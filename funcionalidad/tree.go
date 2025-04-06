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
	//node obtiene el valor al que se le pulso enter
	path := node.GetReference().(string)
	info, err := os.Stat(path) //contiene informacion como si es carpeta o no, peso etc etc
	if err != nil {
		return
	}
	if info.IsDir() {
		//node.GetChildren() arreglo de los elementos que se encuentran dentro de la carpeta
		if len(node.GetChildren()) == 0 {
			helpers.AddChildren(node, path) //se adicionan los elementos de manera grafica
		}
		node.SetExpanded(!node.IsExpanded()) //se expande como si fuera menu desplegable
	} else {
		//si se presiono enter a un archivo entonces se va a abrir en el editor de texto vim
		// se suspende la ejecución del sistema
		app.Suspend(func() {
			cmd := exec.Command("nano", path)
			cmd.Stdin = os.Stdin
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			cmd.Run()
		})
	}
}

func CapturaOpcion(tree *tview.TreeView, app *tview.Application, event *tcell.EventKey, rootNode *tview.TreeNode) *tcell.EventKey {
	node := tree.GetCurrentNode()        //obtenemos el nodo en el que nos encontramos (direccion en memoria)
	path := node.GetReference().(string) //transformamos la direccion en caracteres
	// se leen los caracteres en ascii
	switch event.Rune() {
	case 'c':
		operaciones.Crear(path, tree, app, node)

	case 'e':
		operaciones.Eliminar(path, tree, app, rootNode)

	case 'r':
		operaciones.Renombrar(path, tree, app, rootNode)

	case 'm':
		operaciones.Renombrar(path, tree, app, rootNode)
	}
	return event
}
