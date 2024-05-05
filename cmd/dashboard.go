package cmd

import (
	"task-manager/db"
	"task-manager/utils"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/kancli"
	"github.com/charmbracelet/bubbles/list"
	"github.com/spf13/cobra"
)

// dashboardCmd represents the dashboard command
var dashboardCmd = &cobra.Command{
	Use:   "dashboard",
	Short: "A brief description of your command",
	RunE: func(cmd *cobra.Command, args []string) error {
		t, err := utils.OpenDB(utils.SetupPath())
		if err != nil {
			return err
		}
		defer t.Db.Close()
		todos, err := t.GetTasksByStatus(db.TODO.String())
		if err != nil {
			return err
		}
		inprogress, err := t.GetTasksByStatus(db.INPROGRESS.String())
		if err != nil {
			return err
		}
		done, err := t.GetTasksByStatus(db.DONE.String())
		if err != nil {
			return err
		}

		todoColumn := kancli.NewColumn(tasksToItems(todos), db.TODO, true)
		inprogressColumn := kancli.NewColumn(tasksToItems(inprogress), db.INPROGRESS, false)
		doneColumn := kancli.NewColumn(tasksToItems(done), db.DONE, false)
		board := kancli.NewDefaultBoard([]kancli.Column{todoColumn, inprogressColumn, doneColumn})
		p := tea.NewProgram(board)
		_, err = p.Run()
		return nil
	},
}

func tasksToItems(tasks []db.Task) []list.Item {
	var items []list.Item
	for _, t := range tasks {
		items = append(items, t)
	}
	return items
}

func init() {
	rootCmd.AddCommand(dashboardCmd)
}
