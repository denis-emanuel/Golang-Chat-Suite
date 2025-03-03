/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// userServiceCmd represents the userService command
var userServiceCmd = &cobra.Command{
	Use:   "userService",
	Short: "Start user service",
	Long:  `Start user service`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("userService called")
	},
}

func init() {
	rootCmd.AddCommand(userServiceCmd)
}
