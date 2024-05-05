package cmd

import (
	"fmt"
	"task-manager/db"
	"task-manager/utils"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all tasks",
	RunE: func(cmd *cobra.Command, args []string) error {
		t, err := utils.OpenDB(utils.SetupPath())
		if err != nil {
			return err
		}
		defer t.Db.Close()
		tasks, err := t.GetTasks()
		if err != nil {
			return err
		}
		fmt.Print(setupTable(tasks))
		return nil
	},
}

func setupTable(tasks []db.Task) *table.Table {
	columns := []string{"ID", "Name", "Project", "Status", "Created At"}
	var rows [][]string
	for _, task := range tasks {
		rows = append(rows, []string{
			fmt.Sprintf("%d", task.ID),
			task.Name,
			task.Project,
			task.Status,
			task.Created,
		})
	}
	t := table.New().
		Border(lipgloss.HiddenBorder()).
		Headers(columns...).
		Rows(rows...).
		StyleFunc(func(row, col int) lipgloss.Style {
			if row == 0 {
				return lipgloss.NewStyle().
					Foreground(lipgloss.Color("212")).
					Border(lipgloss.HiddenBorder()).
					BorderTop(false).
					BorderLeft(false).
					BorderRight(false).
					BorderBottom(true).
					Bold(true)
			}
			if row%2 == 0 {
				return lipgloss.NewStyle().Foreground(lipgloss.Color("246"))
			}
			return lipgloss.NewStyle()
		})
	return t
}

func init() {
	rootCmd.AddCommand(listCmd)
}
