package cmd

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use:   "msg",
	Short: "mensaje corto",
	Long:  "mensaje largo",
}

func Execute() {
	cobra.CheckErr((rootCmd.Execute()))
}
