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
	input := tview.NewInputField() //se crea un campo de texto al presionar la opcion correspondiente a la funcion

	input.SetLabel("Nombre nuevo archivo/directorio (termina en / para dir): "). //da la instruccion al usuario
		//funcion que se ejecuta al dar enter
		SetDoneFunc(func(key tcell.Key) {
			name := input.GetText()           //obtiene el texto ingresado por el usuario
			full := filepath.Join(path, name) //se une la ruta en la que se encuentra con el nombre escrito por el usuario
			if strings.HasSuffix(name, "/") {
				os.MkdirAll(full, 0755) // se crea el directorio con los permisos
			} else {
				os.WriteFile(full, []byte(""), 0644) // crea el archivo vacio con los permisos
			}
			helpers.AddChildren(node, path)        //se refresca el nodo en el que se encuentra para que aparezca el archivo o directorio nuevo
			app.SetRoot(tree, true).SetFocus(tree) //vuelve a mostrar el arbol
		})
	app.SetRoot(input, true).SetFocus(input)
}

func Eliminar(path string, tree *tview.TreeView, app *tview.Application, rootNode *tview.TreeNode) {
	modal := tview.NewModal()
	modal.SetText("¿Estás seguro de querer eliminar " + path + "? (S/N)").
		AddButtons([]string{"s", "n"}).
		SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			if strings.ToLower(buttonLabel) == "s" {
				// Eliminar archivo o directorio
				os.RemoveAll(path)

				// Buscar el nodo padre para actualizarlo
				parentPath := filepath.Dir(path)
				parentNode := helpers.FindNodeByPath(rootNode, parentPath)
				if parentNode != nil {
					helpers.AddChildren(parentNode, parentPath)
				}
			}

			// Volver a mostrar el árbol después de la confirmación
			app.SetRoot(tree, true).SetFocus(tree)
		})

	// Mostrar el modal
	app.SetRoot(modal, true).SetFocus(modal)
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
