package operaciones

// https://docs.github.com/en/rest/repos/repos?apiVersion=2022-11-28#list-repositories-for-a-user
// https://docs.github.com/en
// https://labex.io/tutorials/go-how-to-handle-file-read-permissions-450986

import (
	"fmt"
	"os"
	"os/user"
	"strconv"
	"syscall"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func GestionarPermisos(path string, tree *tview.TreeView, app *tview.Application, node *tview.TreeNode) {
	newPrimitive := func(text string) *tview.TextView {
		return tview.NewTextView().
			SetTextAlign(tview.AlignCenter).
			SetText(text)
	}

	fileInfo, err := os.Stat(path)
	if err != nil {
		panic(err)
	}

	stat, ok := fileInfo.Sys().(*syscall.Stat_t)
	if !ok {
		panic("La información del sistema no es de tipo syscall.Stat_t")
	}
	groupId := stat.Gid
	ownerId := stat.Uid

	// Obtenemos permisos en formato string
	mode := fileInfo.Mode()
	permisosAr := mode.String()

	// Obtenemos nombre del grupo
	nombreGrupo, err := obtenerGrupo(int(groupId))
	if err != nil {
		nombreGrupo = "Desconocido"
	}

	// Obtenemos nombre del propietario
	nombrePropietario, err := obtenerUsuario(int(ownerId))
	if err != nil {
		nombrePropietario = "Desconocido"
	}

	nombre := newPrimitive(fmt.Sprintf("Nombre: %s", fileInfo.Name()))
	permisos := newPrimitive(fmt.Sprintf("Permisos: %s", permisosAr))
	grupo := newPrimitive(fmt.Sprintf("Grupo: %s", nombreGrupo))
	propietario := newPrimitive(fmt.Sprintf("Propietario: %s", nombrePropietario))

	grid := tview.NewGrid().
		SetRows(3, 0, 0).
		SetColumns(30, 30, 30, 30).
		SetBorders(true).
		AddItem(newPrimitive("Gestión de Permisos"), 0, 0, 1, 4, 0, 0, false).
		AddItem(nombre, 1, 0, 1, 1, 0, 100, false).
		AddItem(permisos, 1, 1, 1, 1, 0, 100, false).
		AddItem(grupo, 1, 2, 1, 1, 0, 100, false).
		AddItem(propietario, 1, 3, 1, 1, 0, 100, false)

	captura := tview.NewTextView().
		SetDynamicColors(true).
		SetText("Presiona [red]Esc[white] para volver").
		SetTextAlign(tview.AlignCenter)

	captura.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyEsc:
			app.SetRoot(tree, true).SetFocus(tree)
			return nil
		}
		return event
	})

	layout := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(captura, 1, 0, true).
		AddItem(grid, 0, 1, false)

	if err := app.SetRoot(layout, true).SetFocus(captura).Run(); err != nil {
		panic(err)
	}
}

func obtenerGrupo(groupId int) (string, error) {
	gidStr := strconv.Itoa(groupId)
	group, err := user.LookupGroupId(gidStr)
	if err != nil {
		return "", err
	}
	return group.Name, nil
}

func obtenerUsuario(userId int) (string, error) {
	uidStr := strconv.Itoa(userId)
	usuario, err := user.LookupId(uidStr)
	if err != nil {
		return "", err
	}
	return usuario.Username, nil
}
