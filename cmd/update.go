package cmd

import (
	"strconv"
	"task-manager/db"
	"task-manager/utils"
	"time"

	"github.com/spf13/cobra"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "A brief description of your command",
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		t, err := utils.OpenDB(utils.SetupPath())
		if err != nil {
			return err
		}
		defer t.Db.Close()
		name, err := cmd.Flags().GetString("name")
		if err != nil {
			return err
		}
		project, err := cmd.Flags().GetString("project")
		if err != nil {
			return err
		}
		prog, err := cmd.Flags().GetInt("status")
		if err != nil {
			return err
		}
		id, err := strconv.Atoi(args[0])
		if err != nil {
			return err
		}
		var status string
		switch prog {
		case int(db.TODO):
			status = db.TODO.String()
		case int(db.INPROGRESS):
			status = db.INPROGRESS.String()
		case int(db.DONE):
			status = db.DONE.String()
		default:
			status = db.TODO.String()
		}
		newTask := db.Task{
			ID:      id,
			Name:    name,
			Project: project,
			Status:  status,
			Created: time.Now(),
		}
		return t.Update(newTask)
	},
}

func init() {
	updateCmd.Flags().StringP(
		"name",
		"n",
		"",
		"Specify the name for the task",
	)
	updateCmd.Flags().StringP(
		"project",
		"p",
		"",
		"Specify a project for your task",
	)
	updateCmd.Flags().IntP(
		"status",
		"s",
		int(db.TODO),
		"Specify a status for your task",
	)

	rootCmd.AddCommand(updateCmd)
}
