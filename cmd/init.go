/*

Copyright © 2025 Drew Meyaln <dnmeylan@gmail.com>

*/
package cmd

import (
	"fmt"
	"os"
//	"io/fs"
//	"path/filepath"
	"github.com/pterm/pterm"
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

func check(e error) {
	if e != nil {
		panic(e)
	}
}

var initClusterCmd = &cobra.Command{
	Use:	"cluster",
	Short:	"Cluster management",
	Run: func(cmd *cobra.Command, args []string) {
		getClusterName := pterm.DefaultInteractiveTextInput
		getClusterSeed := pterm.DefaultInteractiveTextInput
		getClusterKey := pterm.DefaultInteractiveTextInput
		getClusterName.DefaultText = "What is the name of the cluster you are initializing?"
		getClusterSeed.DefaultText = "Enter a local IP address to serve as the cluster seed IP: " 
		getClusterKey.DefaultText = "Enter the filepath for the gpg key: "
		Name, _ := getClusterName.Show()
		Seed, _ := getClusterSeed.Show()
		Key, _ := getClusterKey.Show()

		fmt.Println("Name: ", Name)
		fmt.Println("Seed: ", Seed)
		fmt.Println("Key: ", Key)
	},
}

var initLocalCmd = &cobra.Command {
	Use: 	"local",
	Short:	"Local management",
	Run: func(cmd *cobra.Command, args []string) {
		getDBName := pterm.DefaultInteractiveTextInput
		getDBKey := pterm.DefaultInteractiveTextInput
		getDBName.DefaultText = "Enter a local database name: " 
		getDBKey.DefaultText = "Would you like to use an existing gpg key for encryption? (Y/N) [n]"
		getGPGKeyPass.DefaultText = "Enter a password to use for this stores gpg key: "
		getGPGKeyPass2.DefaultText = "Enter your password again: "

		Name, _ := getDBName.Show()
		Key, _ := getDBKey.Show()
		if Key != 

		err := os.Mkdir("./internal/"+Name, 0755) // creates the directory to store the passwords.
		check(err) // throws panic if previous command fails

		// fmt.Println("Name: ", Name)
		// fmt.Println("Key: ", Key)
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
