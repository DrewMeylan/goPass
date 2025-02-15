/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("create called")
	},
}

var createLocalCmd = &cobra.Command{
	Use: "local",
	Short: "Local management",
	Run: func(cmd *cobra.Command, args []string) {
		getDBName := pterm.DefaultInteractiveTextInput
		getDBName.DefaultText = "What would you like to name this local database?"
		result, _ := getDBName.Show()
		fmt.Println(result)
	},
}

var createClusterCmd = &cobra.Command{
	Use: "cluster",
	Short: "Cluster management",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("create cluster called")
	},
}


func init() {
	rootCmd.AddCommand(createCmd)
	createCmd.AddCommand(createLocalCmd, createClusterCmd)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
