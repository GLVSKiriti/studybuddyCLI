package cmd

import (
	"github.com/GLVSKiriti/studybuddy/data"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "To see the list of your notes at once",
	Long:  `To see the list of your notes at once`,
	Run: func(cmd *cobra.Command, args []string) {
		data.DisplayAllNotes()
	},
}

func init() {
	noteCmd.AddCommand(listCmd)
}
