package cmd

import (
	"github.com/spf13/cobra"
	"github.com/stuartskelton/todolist/todolist"
	"strings"
)

var deleteCmd = &cobra.Command{
	Use:     "delete",
	Aliases: []string{"d"},
	Short:   "Delete a todo by its Id",
	Long: `Delete a todo by its Id
todo d 33
will delete a todo with id 33`,
	Run: func(cmd *cobra.Command, args []string) {
		todolist.NewApp().DeleteTodo(strings.Join(args, " "))
	},
}

var completeCmd = &cobra.Command{
	Use:     "complete",
	Aliases: []string{"c"},
	Short:   "Complete (c) a todo by its Id",
	Long: `Complete (c) a todo by its Id
todo c 33
Completes a todo with id 33`,
	Run: func(cmd *cobra.Command, args []string) {
		todolist.NewApp().CompleteTodo(strings.Join(args, " "))
	},
}

var uncompleteCmd = &cobra.Command{
	Use:     "uncomplete",
	Aliases: []string{"uc"},
	Short:   "Uncomplete (uc) a todo by its Id",
	Long: `Uncomplete (uc) a todo by its Id
todo uc 33
Uncompletes a todo with id 33`,
	Run: func(cmd *cobra.Command, args []string) {
		todolist.NewApp().UncompleteTodo(strings.Join(args, " "))
	},
}

var archiveCmd = &cobra.Command{
	Use:     "archive",
	Aliases: []string{"ar"},
	Short:   "Archive (ar) a todo by its Id",
	Long: `Archive (ar) a todo by its Id
todo ar 33
Archives a todo with id 33`,
	Run: func(cmd *cobra.Command, args []string) {
		todolist.NewApp().ArchiveTodo(strings.Join(args, " "))
	},
}

var unarchiveCmd = &cobra.Command{
	Use:     "unarchive",
	Aliases: []string{"uar"},
	Short:   "Unarchive (uar) a todo by its Id",
	Long: `Unarchive (uar) a todo by its Id
todo uar 33
Unarchives a todo with id 33`,
	Run: func(cmd *cobra.Command, args []string) {
		todolist.NewApp().UnarchiveTodo(strings.Join(args, " "))
	},
}

var prioritizeCmd = &cobra.Command{
	Use:     "prioritize",
	Aliases: []string{"p"},
	Short:   "Psrioritize (p) a todo by its Id",
	Long: `Prioritize (p) a todo by its Id
todo p 33
Prioritizes a todo with id 33`,
	Run: func(cmd *cobra.Command, args []string) {
		todolist.NewApp().PrioritizeTodo(strings.Join(args, " "))
	},
}

var unprioritizeCmd = &cobra.Command{
	Use:     "unprioritize",
	Aliases: []string{"up"},
	Short:   "Unprioritize (up) a todo by its Id",
	Long: `Unprioritize (up) a todo by its Id
todo up 33
Unprioritizes a todo with id 33`,
	Run: func(cmd *cobra.Command, args []string) {
		todolist.NewApp().UnprioritizeTodo(strings.Join(args, " "))
	},
}

func init() {
	// delete a task
	RootCmd.AddCommand(deleteCmd)

	// complete a task
	RootCmd.AddCommand(completeCmd)

	// uncomplete a task
	RootCmd.AddCommand(uncompleteCmd)

	// archive a task
	RootCmd.AddCommand(archiveCmd)

	// unarchive a task
	RootCmd.AddCommand(unarchiveCmd)

	// prioritize a task
	RootCmd.AddCommand(prioritizeCmd)

	// unprioritize a task
	RootCmd.AddCommand(unprioritizeCmd)
}
