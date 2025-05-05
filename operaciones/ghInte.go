package operaciones

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"sistemagestoarchivos/helpers"

	"github.com/gdamore/tcell/v2"
	"github.com/google/go-github/github"
	"github.com/rivo/tview"
)

func ObtenerUsuarioRepo(path string, tree *tview.TreeView, app *tview.Application, node *tview.TreeNode) {
	input := tview.NewInputField() //se crea el input para ingresar el nombre del usuario
	input.SetLabel("Ingrese el nombre de usuario: ").
		SetDoneFunc(func(key tcell.Key) { //lo que se va a realizar al presionar enter
			nombreUsuario := input.GetText()              //obtenemos la entrada
			repos := ObtencionRepositorios(nombreUsuario) //buscamos los repositorios del usuario
			list := tview.NewList()                       //creamos una lista en donde se colocara los repositorios
			list.SetBorder(true).SetTitle("Repos de " + nombreUsuario)
			for _, repo := range repos {
				if repo.Name != nil {
					//AddItem(texto principal, texto secundario o descripcion,boton que seleccionara, funcion que se ejecutara al seleccionar)
					nombreRepo := *repo.Name
					urlRepo := *repo.CloneURL
					list.AddItem(nombreRepo, "", 0, func() {
						ClonarRepo(path, nombreRepo, urlRepo, node)
					}) //se añaden los repositorios a la lista
				}
			}

			// Agregamos una opción para volver o salir
			list.AddItem("Salir", "Cerrar la app", 0, func() {
				app.SetRoot(tree, true).SetFocus(tree) //cambio de pantalla
			})
			app.SetRoot(list, true).SetFocus(list)
		})
	app.SetRoot(input, true).SetFocus(input)
}

// org es el nombre del usuario de github
func ObtencionRepositorios(org string) []*github.Repository {
	ctx := context.Background()     //encargado de la concurrencia para la optimizacion al realizar la peticion de los datos
	client := github.NewClient(nil) //construye un cliente de github para tener acceso a las diferentes partes de la api de github

	//formato para la peticion de repositorios
	opt := &github.RepositoryListOptions{
		ListOptions: github.ListOptions{PerPage: 10},
	}

	//arreglo en el cual se almacenaran todos los repositorios pedidos
	var allRepos []*github.Repository

	// bucle infinito, es lo mismo que for true
	for {
		repos, resp, err := client.Repositories.List(ctx, org, opt) // retorno de los repositorios
		if err != nil {
			fmt.Println("Error:", err)
			break
		}

		allRepos = append(allRepos, repos...) //almacenamiento de los repositorios en el arreglo

		if resp.NextPage == 0 {
			break
		}
		opt.Page = resp.NextPage
	}

	return allRepos
}

func ClonarRepo(path string, repositorio string, urlRepo string, node *tview.TreeNode) {
	cmd := exec.Command("git", "clone", urlRepo)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	helpers.AddChildren(node, path) //refresca el arbol
	if err != nil {
		fmt.Println("Error al clonar el repositorio:", err)
	}
}
