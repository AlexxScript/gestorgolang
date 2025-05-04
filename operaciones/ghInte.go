package operaciones

import (
	"context"
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/google/go-github/github"
	"github.com/rivo/tview"
)

func ObtenerUsuarioRepo(path string, tree *tview.TreeView, app *tview.Application, node *tview.TreeNode) {
	input := tview.NewInputField()
	input.SetLabel("Ingrese el nombre de usuario: ").
		SetDoneFunc(func(key tcell.Key) {
			nombreUsuario := input.GetText()
			repos := ObtencionRepositorios(nombreUsuario)
			list := tview.NewList()
			list.SetBorder(true).SetTitle("Repos de " + nombreUsuario)
			for _, repo := range repos {
				if repo.Name != nil {
					list.AddItem(*repo.Name, "", 0, nil)
				}
			}

			// Agregamos una opción para volver o salir
			list.AddItem("Salir", "Cerrar la app", 0, func() {
				app.SetRoot(tree, true).SetFocus(tree)
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

	// Imprimir los nombres
	// for _, repo := range allRepos {
	// 	fmt.Println(*repo.Name)
	// }
	return allRepos
}
