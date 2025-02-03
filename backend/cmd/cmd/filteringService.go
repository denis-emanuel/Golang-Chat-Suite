/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// filteringServiceCmd represents the filteringService command
var filteringServiceCmd = &cobra.Command{
	Use:   "filteringService",
	Short: "Start the chat filtering service",
	Long:  "Start the chat filtering service",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("filteringService called")
	},
}

func init() {
	rootCmd.AddCommand(filteringServiceCmd)
}
