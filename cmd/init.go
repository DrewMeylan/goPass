/*

Copyright © 2025 Drew Meyaln <dnmeylan@gmail.com>

*/
package cmd

import (
	"fmt"
	"strings"
	"os"
//	"io/fs"
//	"path/filepath"
	"syscall"
	"gopass/intenral/crypto"
  "gopass/internal/password"
  "gopass/internal/files"
  "golang.org/x/term"
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
// ------------------------------------------------ Staging -----------------------------------------------------
		getDBName := pterm.DefaultInteractiveTextInput
		getDBKey := pterm.DefaultInteractiveTextInput
		getRecipient := pterm.DefaultInteractiveTextInput
    getUserName := pterm.DefaultInteractiveTextInput
    getDBName.DefaultText = "Enter a local database name: "
    getDBKey.DefaultText = "Would you like to use an existing gpg key for encryption? (Y/N) [n] "
    // If no
    getUserName.DefaultText = "Enter your name: "
    getRecipient.DefaultText = "Enter your e-mail: "
// ------------------------------------------------- Logic ------------------------------------------------------	
		var gpgPassword string // I'm not certain this is necessary?

    Name, _ := getDBName.Show()
		useExistKey, _ := getDBKey.Show() // Maybe offload this to helper function like getPass()
		useExistNorm := strings.ToLower(strings.TrimSpace(useExistKey)) // With above
		if useExistNorm == "yes" || useExistNorm == "y" {
			existKey, _ := getRecipient.Show()
			fmt.Println("Creating database " + Name + " with key for " + existKey) 
      // Modify config file --> WriteFile() function in internal/files
      // create directory --> CreateFolder() function in internal/files?
		} else {
      name, _ = getUserName.show()
      email, _ = getRecipient.show()
			gpgPassword, err = GetPassword()
			fmt.Println("Attempted to generate GPG key... ")
      _, err := GenerateGPGKey(name, email, gpgPassword)
      if err != nil {
        fmt.Println("Error encountered: " + err)
      }
      else {
        fmt.Println("GPG key created successfully")
      }
      // Modify config file
      // create directory
		}

		err := os.Mkdir("./internal/"+Name, 0755) // creates the directory to store the passwords. // THIS SHOULD BE OFF-LOADED INTO THE FILES PKG
		check(err) // throws panic if previous command fails

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
