package cmd

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use:   "atajos",
	Short: "[c] Crear   [d] Eliminar   [r] Renombrar   [m] Mover",
	Long:  "[c] Crear   [d] Eliminar   [r] Renombrar   [m] Movermensaje",
}

func Execute() {
	cobra.CheckErr((rootCmd.Execute()))
}
