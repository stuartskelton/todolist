package cmd

import (
	"strings"

	"github.com/spf13/cobra"
	"github.com/stuartskelton/todolist/todolist"
)

var expandCmd = &cobra.Command{
	Use:     "expand",
	Aliases: []string{"ex"},
	Short:   "Expand (ex) an existing todo.",
	Long: `Expand (ex) an existing todo.
todo ex 39 +final: read physics due mon, do literature report due fri
Removes the todo with id 39, and adds following two todos`,
	Run: func(cmd *cobra.Command, args []string) {
		todolist.NewApp().ExpandTodo(strings.Join(args, " "))
	},
}

func init() {
	// expand a todos
	RootCmd.AddCommand(expandCmd)
}
