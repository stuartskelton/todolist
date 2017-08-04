package cmd

import (
	"github.com/spf13/cobra"
	"github.com/stuartskelton/todolist/todolist"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize your todos repo",
	Long:  `Initialize your todos repo, named .todos.json`,
	Run: func(cmd *cobra.Command, args []string) {
		todolist.NewApp().InitializeRepo()
	},
}

// archive any completed todo item
var archiveCompletedCmd = &cobra.Command{
	Use:   "ac",
	Short: "Archives all completed todos",
	Long: `Archives all completed todos. By archiving
your completed todos these will be hidden by default when
listing your current Todos.

To view them again you will have to 'todolist list archived'.`,
	Run: func(cmd *cobra.Command, args []string) {
		todolist.NewApp().ArchiveCompleted()
	},
}

// Collect and delete the garbage in the todo repo
var garbageCollectCmd = &cobra.Command{
	Use:   "gc",
	Short: "Deletes all archived todos.",
	Long: `Deletes all archived todos. You will not be able to
undo this command, unless you have a backup.`,
	Run: func(cmd *cobra.Command, args []string) {
		todolist.NewApp().GarbageCollect()
	},
}

func init() {
	// init a todo repo
	RootCmd.AddCommand(initCmd)

	// archive all the comepleted todos
	RootCmd.AddCommand(archiveCompletedCmd)

	// run the todo garbage collection
	RootCmd.AddCommand(garbageCollectCmd)
}
