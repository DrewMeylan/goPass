/*

Copyright © 2025 Drew Meyaln <dnmeylan@gmail.com>

*/
package cmd

import (
	"fmt"
	"os"
//	"io/fs"
//	"path/filepath"
	"syscall"
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

// OFFLOAD TO HELPER PACKAGE
func check(e error) {
	if e != nil {
		panic(e)
	}
}

// getPassword prompts the user twice for a password using pterm and ensures they match.
// OFFLOAD TO HELPER PACKAGE
func getPassword() string {
	for {
		// Prompt for first password
		pterm.Info.Println("Enter your password:")
		password1, err := term.ReadPassword(int(syscall.Stdin))
		fmt.Println() // Move to a new line after input
		if err != nil {
			pterm.Error.Println("Error reading password:", err)
			continue
		}

		// Prompt for confirmation
		pterm.Info.Println("Re-enter your password:")
		password2, err := term.ReadPassword(int(syscall.Stdin))
		fmt.Println()
		if err != nil {
			pterm.Error.Println("Error reading password:", err)
			continue
		}

		// Check if passwords match
		if string(password1) == string(password2) {
			pterm.Success.Println("Passwords match!")
			return string(password1)
		} else {
			pterm.Warning.Println("Passwords do not match. Please try again.")
		}
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
// ------------------------------------------------ Staging -----------------------------------------------------
		getDBName := pterm.DefaultInteractiveTextInput
		getDBKey := pterm.DefaultInteractiveTextInput
		getKeyPath := pterm.DefaultInteractiveTextInput
		getKeyPath.DefaultText = "Enter the file path for your gpg key: "
		getDBName.DefaultText = "Enter a local database name: " 
		getDBKey.DefaultText = "Would you like to use an existing gpg key for encryption? (Y/N) [n] "
// ------------------------------------------------- Logic ------------------------------------------------------	
		Name, _ := getDBName.Show()
		useExistKey, _ := getDBKey.Show() // Maybe offload this to helper function like getPass()
		useExistNorm := strings.ToLower(string.TrimSpace(useExistKey)) // With above
		var gpgPassword string
		if useExistNorm == "yes" || useExistNorm == "y" {
			existKeyPath, _ := getKeyPath.Show()
			fmt.Println("Creating database " + Name + " with key " + existKeyPath) 
		} else {
			gpgPassword = getPassword()
			fmt.Println("password" gpgPassword)
		}

		err := os.Mkdir("./internal/"+Name, 0755) // creates the directory to store the passwords.
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
