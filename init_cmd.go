package main

// initCmd represents the init command
import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/ubclaunchpad/inertia/client"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize an inertia project in this repository",
	Long: `Initialize an inertia project in this GitHub repository.
There must be a local git repository in order for initialization
to succeed.`,
	Run: func(cmd *cobra.Command, args []string) {
		version := cmd.Parent().Version
		givenVersion, _ := cmd.Flags().GetString("version")
		if givenVersion != version {
			version = givenVersion
		}

		err := client.InitializeInertiaProject(cmd.Parent().Version)
		if err != nil {
			log.Fatal(err)
		}
		println("A .inertia.toml configuration file has been created to store")
		println("Inertia configuration. It is recommended that you DO NOT commit")
		println("this file in source control since it will be used to store")
		println("sensitive information.")
		println("\nYou can now use 'inertia remote add' to connect your remote")
		println("VPS instance.")
	},
}

var resetCmd = &cobra.Command{
	Use:   "reset",
	Short: "Reset the Inertia project in this repository.",
	Long: `Reset removes the existing Inertia configuration from
	this repository.`,
	Run: func(cmd *cobra.Command, args []string) {
		println("WARNING: This will remove your current Inertia configuration")
		println("and is irreversible. Continue? (y/n)")
		var response string
		_, err := fmt.Scanln(&response)
		if err != nil {
			log.Fatal("invalid response - aborting")
		}
		if response != "y" {
			log.Fatal("aborting")
		}
		path, err := client.GetConfigFilePath()
		if err != nil {
			log.Fatal(err)
		}
		os.Remove(path)
		println("Inertia configuration removed.")
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
	rootCmd.AddCommand(resetCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	addCmd.Flags().String("version", "default", "Inertia daemon version")
}