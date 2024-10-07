package cmd

import (
    "os"
    "os/exec"
    "fmt"
    "github.com/spf13/cobra"
)

var precommitSetupCmd = &cobra.Command{
    Use:   "precommit-setup",
    Short: "Sets up pre-commit hooks for code quality and security checks",
    Long:  `This command sets up pre-commit hooks to automatically run tools like Black, Flake8, and Bandit before commits.`,
    Run: func(cmd *cobra.Command, args []string) {
        // Step 1: Generate the .pre-commit-config.yaml file
        configContent := `
repos:
- repo: https://github.com/psf/black
  rev: stable
  hooks:
  - id: black

- repo: https://github.com/PyCQA/bandit
  rev: stable
  hooks:
  - id: bandit
`
        err := os.WriteFile(".pre-commit-config.yaml", []byte(configContent), 0644)
        if err != nil {
            fmt.Println("Error creating pre-commit configuration file:", err)
            return
        }
        fmt.Println("Created .pre-commit-config.yaml file.")

        // Install the pre-commit tool if it's not installed
        cmdCheck := exec.Command("cmd", "/C", "env\\Scripts\\pre-commit.exe", "--version")
        err = cmdCheck.Run()
        if err != nil {
            fmt.Println("pre-commit is not installed. Installing it now...")
            cmdInstall := exec.Command("cmd", "/C", "env\\Scripts\\pip.exe", "install", "pre-commit")
            err = cmdInstall.Run()
            if err != nil {
                fmt.Println("Error installing pre-commit:", err)
                return
            }
            fmt.Println("Installed pre-commit.")
        }

        // Run 'pre-commit install' to install the hooks
        cmdInstallHooks := exec.Command("cmd", "/C", "env\\Scripts\\pre-commit.exe", "install")
        err = cmdInstallHooks.Run()
        if err != nil {
            fmt.Println("Error installing pre-commit hooks:", err)
            return
        }
        fmt.Println("Pre-commit hooks installed successfully.")
    },
}