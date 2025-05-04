package operaciones

import (
	"context"
	"fmt"

	"github.com/google/go-github/github"
)

// org es el nombre del usuario de github
func ListarRepositorios(org string) {
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
	for _, repo := range allRepos {
		fmt.Println(*repo.Name)
	}
}
