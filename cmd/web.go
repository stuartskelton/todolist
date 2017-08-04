package cmd

import (
	"fmt"
	"os"

	"github.com/skratchdot/open-golang/open"
	"github.com/spf13/cobra"
	"github.com/stuartskelton/todolist/todolist"
)

// webCmd represents the web command
var webCmd = &cobra.Command{
	Use:   "web",
	Short: "Run a web interface interact with your Todos.",
	Long: `Run a web interface interact with your Todos.
when running go to http://localhost:7890 to see and interacti with your
Todo list.`,
	Run: func(cmd *cobra.Command, args []string) {
		app := todolist.NewApp()

		if err := app.Load(); err != nil {
			os.Exit(1)
		} else {
			web := todolist.NewWebapp()
			fmt.Println("Now serving todolist web.\nHead to http://localhost:7890 to see your todo list!")
			open.Start("http://localhost:7890")
			web.Run()
		}
	},
}

func init() {
	// run a tiny webserver for the todos
	RootCmd.AddCommand(webCmd)
}
