package cmd

import (
	"strings"

	"github.com/spf13/cobra"
	"github.com/stuartskelton/todolist/todolist"
)

var editCmd = &cobra.Command{
	Use:     "edit",
	Aliases: []string{"e"},
	Short:   "edit (e) a todos due dates",
	Long: `Edit (e) a todos due dates
todo e 33 due mon
Edits the todo with 33 and sets the due date to this coming Monday`,
	Run: func(cmd *cobra.Command, args []string) {
		todolist.NewApp().EditTodo(strings.Join(args, " "))
	},
}

func init() {
	// edit a todos
	RootCmd.AddCommand(editCmd)
}
