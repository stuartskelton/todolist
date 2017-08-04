package cmd

import (
	"strings"

	"github.com/spf13/cobra"
	"github.com/stuartskelton/todolist/todolist"
)

var listCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"l"},
	Short:   "list your todos",
	Long: `When listing todos, you can filter and group the output.

todo l due (tod|today|tom|tomorrow|overdue|this week|next week|last week|mon|tue|wed|thu|fri|sat|sun|none)
todo l overdue
Filtering by date:

todo l due tod
lists all todos due today

todo l due tom
lists all todos due tomorrow

todo l due mon
lists all todos due monday

todo l overdue
lists all todos where the due date is in the past

todo l completed (tod|today|this week)
Filtering by date:

todo l completed (tod|today)
lists all todos that were completed today

todo l completed this week
lists all todos that were completed this week`,
	Run: func(cmd *cobra.Command, args []string) {
		todolist.NewApp().ListTodos(strings.Join(args, " "))
	},
}

var agendaCmd = &cobra.Command{
	Use:   "agenda",
	Short: "Lists all todos where the due date is today or in the past",
	Long: `todo agenda
lists all todos where the due date is today or in the past`,
	Run: func(cmd *cobra.Command, args []string) {
		todolist.NewApp().ListTodos("agenda")
	},
}

func init() {
	// list the todos
	RootCmd.AddCommand(listCmd)
	// shows the agenda
	RootCmd.AddCommand(agendaCmd)
}
