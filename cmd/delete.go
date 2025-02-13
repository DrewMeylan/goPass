/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("delete called")
	},
}

var deleteClusterCmd = &cobra.Command{
	Use: "cluster",
	Short: "Cluster deletion",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("delete cluster called")
	},
}

var deleteLocalCmd = &cobra.Command{
	Use: "local",
	Short: "local deletion",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("delete local called")
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
	deleteCmd.AddCommand(deleteLocalCmd, deleteClusterCmd)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
