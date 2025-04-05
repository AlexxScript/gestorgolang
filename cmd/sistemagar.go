package cmd

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"github.com/spf13/cobra"

	"sistemagestoarchivos/funcionalidad"
)

var listUICmd = &cobra.Command{
	Use:   "list-ui [directory]", //nombre del comando con el cual se ejecutara la funcionalidad
	Short: "Explorador interactivo de archivos",
	Args:  cobra.MaximumNArgs(1),
	Run:   showfilesanddirectories, //funcion que se ejecutara al ejecutar el comando
}

func init() {
	rootCmd.AddCommand(listUICmd) //inicializacion del comando
}

// cmd puntero hacia el paquete decobra
// argumentos que se van a pasar
func showfilesanddirectories(cmd *cobra.Command, args []string) {
	app := tview.NewApplication() //creacion del cli intercativo
	rootDir := "."

	if len(args) > 0 {
		rootDir = args[0]
	}

	rootNode := tview.NewTreeNode(rootDir).SetReference(rootDir).SetExpanded(true)
	tree := tview.NewTreeView().SetRoot(rootNode).SetCurrentNode(rootNode)

	//funcion anonima que enlista los directorios y archivos en forma de arbol
	funcionalidad.AddChildren(rootNode, rootDir)

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
