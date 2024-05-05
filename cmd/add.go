/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"task-manager/utils"
	"github.com/spf13/cobra"
	_ "github.com/mattn/go-sqlite3"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new task",
	RunE: func(cmd *cobra.Command, args []string) error {
		t, err := utils.OpenDB(utils.SetupPath())
		if err != nil {
			return err
		}
		defer t.Db.Close()
		project, err := cmd.Flags().GetString("project")
		if err != nil {
			return err
		}
		if err := t.Insert(args[0], project); err != nil {
			return err
		}
		return nil
	},
}

func init() {
	addCmd.Flags().StringP(
		"project",
		"p",
		"",
		"Specify the project for the task",
	)
	rootCmd.AddCommand(addCmd)
}
