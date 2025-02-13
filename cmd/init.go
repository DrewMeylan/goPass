/*

Copyright © 2025 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a new password store.",
	Run: func(cmd *cobra.Command, args []string) {
		toggle, _ := cmd.Flags().GetBool("toggle")
		if toggle {
			fmt.Println("Help goes here")
		} else {
		fmt.Println("init called")
	}
	},
}

var initClusterCmd = &cobra.Command{
	Use:	"cluster",
	Short:	"Cluster management",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Init cluster called")
	},
}

var initLocalCmd = &cobra.Command {
	Use: 	"local",
	Short:	"Local management",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("init local called")
	},
}	



func init() {
	rootCmd.AddCommand(initCmd)
	initCmd.AddCommand(initClusterCmd, initLocalCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
