/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
//	"gopass/internal/"
	"fmt"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("INSERT CREATION SUBCOMMAND HELP HERE")
	},
}

var createLocalCmd = &cobra.Command{
	Use: "local",
	Short: "Local management",
	Run: func(cmd *cobra.Command, args []string) error { // 
//--- Staging		
		getFileName := pterm.DefaultInteractiveTextInput
		getRecipient := pterm.DefaultInteractiveTextInput // Try and remove eventually; use db ctx
		getFileName.DefaultText = "Name of account: "
		getRecipient.DefaultText = "Enter the associated with the gpg key: "
//--- Collection
		filename := getFileName.show()
		recipient := getRecipient.show()
		password := utils.GetPassword()
//--- Logic
		err := utils.WriteFile(filename, recipient, password) // UPDATE TO REFLECT NEW FILE STRUCTURE
		if err != nil {
			fmt.Println("Error encountered creating password: ", err )
			return err
		}
		else return nil
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
