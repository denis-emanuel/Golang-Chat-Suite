/*
Copyright © 2025 Denis Pop
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "cmd",
	Short: "Chatting and filtering app",
	Long:  `Chatting and filtering app`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
}

func initConfig() {
	// Load .env file manually
	if err := godotenv.Load(".env"); err != nil {
		fmt.Println("No .env file found, using system environment variables.")
	}

	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}

	_ = viper.BindEnv("TEST")

	var cfg config
	if err := viper.Unmarshal(&cfg); err != nil {
		fmt.Println(err)
	}
	fmt.Println(cfg.Test)
}

type config struct {
	Test string `mapstructure:"TEST"`
}
