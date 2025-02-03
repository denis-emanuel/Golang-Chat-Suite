/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// chatServiceCmd represents the chatService command
var chatServiceCmd = &cobra.Command{
	Use:   "chatService",
	Short: "Start the chat service",
	Long:  "Start the chat service",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("chatService called")
	},
}

func init() {
	rootCmd.AddCommand(chatServiceCmd)
}
