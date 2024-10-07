package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"os"

    "github.com/yahialm/django-simplifier/pkg"
)

var rootCmd = &cobra.Command{
	Use:   "dm",
	Short: "A CLI tool to simplify Django project management",
	Long:  `DjangoMate is a CLI tool that streamlines Django project setup and management.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	// Add subcommands here
	rootCmd.AddCommand(serverCmd) // serve
	rootCmd.AddCommand(setupCmd)  // setup
	rootCmd.AddCommand(appCmd)    // Create app - add app to INSTALLED_APPS in setting.py
	rootCmd.AddCommand(migCmd)    // migrate
	rootCmd.AddCommand(mkmigCmd)  // makemigrations
	rootCmd.AddCommand(shmigCmd)  // showmigrations
	rootCmd.AddCommand(dbshCmd)   // dbshell
	rootCmd.AddCommand(shCmd)     // shell
	rootCmd.AddCommand(testCmd)   // test
	rootCmd.AddCommand(suCmd)     // createsuperuser
    rootCmd.AddCommand(pkg.ScanCmd) // Scan
    rootCmd.AddCommand(pkg.FormatCmd) // Format code
    rootCmd.AddCommand(precommitSetupCmd)  // CI & Best practices
}
