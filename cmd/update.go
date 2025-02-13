/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("update called")
	},
}

var updateLocalCmd = &cobra.Command{
	Use: "local",
	Short: "local updater",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("update local called")
	},
}

var updateClusterCmd = &cobra.Command{
	Use: "cluster",
	Short: "cluster updater",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("update cluster called")
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
	updateCmd.AddCommand(updateLocalCmd, updateClusterCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// updateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// updateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
