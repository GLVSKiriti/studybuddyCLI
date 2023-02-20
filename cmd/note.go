package cmd

import (
	"github.com/spf13/cobra"
)

var noteCmd = &cobra.Command{
	Use:   "note",
	Short: "A note command is used to add a new note",
	Long:  `A note command is used to add a new note`,
}

func init() {
	rootCmd.AddCommand(noteCmd)
}
