package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "solopg",
	Short: "SoloPG CLI",
	Long:  "SoloPG is a CLI to manage a solo RPG session in the terminal.",
	Example: "solopg run\n" +
		"solopg load data/template/characters/hero.yaml\n" +
		"solopg version",

	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
