package cmd

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"github.com/spf13/cobra"

	"sistemagestoarchivos/funcionalidad"
	"sistemagestoarchivos/helpers"
)

var sgest = &cobra.Command{
	Use:   "sgest [directory]", //nombre del comando con el cual se ejecutara la funcionalidad
	Short: "Explorador interactivo de archivos",
	Args:  cobra.MaximumNArgs(1),
	Run:   showfilesanddirectories, //funcion que se ejecutara al ejecutar el comando
}

func init() {
	rootCmd.AddCommand(sgest) //inicializacion del comando
}

// cmd puntero hacia el paquete decobra
// argumentos que se van a pasar
func showfilesanddirectories(cmd *cobra.Command, args []string) {
	// caja := tview.NewBox().SetBorder(true).SetTitle("Gestor de archivos en golang")
	app := tview.NewApplication() //creacion del cli intercativo
	rootDir := "."

	if len(args) > 0 {
		rootDir = args[0]
	}

	rootNode := tview.NewTreeNode(rootDir).SetReference(rootDir).SetExpanded(true)
	tree := tview.NewTreeView().SetRoot(rootNode).SetCurrentNode(rootNode)
	tree.SetBorder(true).SetTitle(" Explorador 📁 | 'c' Crear   'e' Eliminar   'r' Renombrar   'm' Mover")

	//funcion anonima que enlista los directorios y archivos en forma de arbol
	helpers.AddChildren(rootNode, rootDir)

	//lo que se ejecutara cuando presiones enter
	tree.SetSelectedFunc(func(node *tview.TreeNode) {
		funcionalidad.ControlarEnter(node, app)
	})

	//captura las teclas que se presionan al momento de seleccionar capturar o remover
	tree.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		return funcionalidad.CapturaOpcion(tree, app, event, rootNode)
	})

	app.SetRoot(tree, true).EnableMouse(true)
	if err := app.Run(); err != nil {
		panic(err)
	}
}
