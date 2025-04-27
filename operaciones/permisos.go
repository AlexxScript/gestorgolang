package operaciones

import (
	"github.com/rivo/tview"
)

func GestionarPermisos(path string, tree *tview.TreeView, app *tview.Application, node *tview.TreeNode) {
	newPrimitive := func(text string) tview.Primitive {
		return tview.NewTextView().
			SetTextAlign(tview.AlignCenter).
			SetText(text)
	}

	nombre := newPrimitive("Nombre del archivo")
	permisos := newPrimitive("Permisos del archivo")
	grupo := newPrimitive("Grupo al que pertenece el archivo")
	propietario := newPrimitive("Propietario del archivo")
	grid := tview.NewGrid().
		SetRows(3, 0).              //dos filas
		SetColumns(30, 30, 30, 30). //4 columnas
		SetBorders(true).

		//AddItem(elemento, fila, columna, cuantasFilasOcupara, columnasAOcupar, alturaMinima, anchoMinimo, focus)
		AddItem(newPrimitive("Gestion de permisos"), 0, 0, 1, 4, 0, 0, false).
		AddItem(nombre, 1, 0, 1, 1, 0, 100, false).
		AddItem(permisos, 1, 1, 1, 1, 0, 100, false).
		AddItem(grupo, 1, 2, 1, 1, 0, 100, false).
		AddItem(propietario, 1, 3, 1, 1, 0, 100, false)

	if err := tview.NewApplication().SetRoot(grid, true).SetFocus(grid).Run(); err != nil {
		panic(err)
	}
}
