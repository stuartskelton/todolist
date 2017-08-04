package cmd

import (
	"strings"

	"github.com/spf13/cobra"
	"github.com/stuartskelton/todolist/todolist"
)

var addCmd = &cobra.Command{
	Use:     "add",
	Aliases: []string{"a"},
	Short:   "Add (a) a todo",
	Long: `Add (a) a todo
the 'a' command adds todos.
You can also optionally specify a due date.
Specify a due date by putting 'due <date>' at the end, where <date> is in (tod|today|tom|tomorrow|mon|tue|wed|thu|fri|sat|sun)

Examples for adding a todo:
todo a Meeting with @bob about +importantPrject due today
todo a +work +verify did @john fix the build\?`,
	Run: func(cmd *cobra.Command, args []string) {
		todolist.NewApp().AddTodo(strings.Join(args, " "))
	},
}

func init() {
	// add a todos
	RootCmd.AddCommand(addCmd)
}
