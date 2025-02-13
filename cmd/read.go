/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// readCmd represents the read command
var readCmd = &cobra.Command{
	Use:   "read",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("read called")
	},
}

var readClusterCmd = &cobra.Command{
	Use: "cluster",
	Short: "Cluster management",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("read cluster called")
	},
}

var readLocalCmd = &cobra.Command{
	Use: "local",
	Short: "local management",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("read local called")
	},
}

func init() {
	rootCmd.AddCommand(readCmd)
	readCmd.AddCommand(readLocalCmd, readClusterCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// readCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// readCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
