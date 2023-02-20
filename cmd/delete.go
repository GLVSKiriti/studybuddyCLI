package cmd

import (
	"github.com/GLVSKiriti/studybuddy/data"
	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "This command is useful to delete a note",
	Long:  `This command is useful to delete a note`,
	Run: func(cmd *cobra.Command, args []string) {
		SelectAndDeleteNote()
	},
}

func init() {
	noteCmd.AddCommand(deleteCmd)
}

func SelectAndDeleteNote() {
	data.DeleteNote()
}
