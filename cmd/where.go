/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"task-manager/utils"

	"github.com/spf13/cobra"
)

// whereCmd represents the where command
var whereCmd = &cobra.Command{
	Use:   "where",
	Short: "Show where your tasks are stored",
	RunE: func(cmd *cobra.Command, args []string) error {
		_, err := fmt.Println("Your tasks are stored in: ", utils.SetupPath())
		if err != nil {
			return err
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(whereCmd)
}
