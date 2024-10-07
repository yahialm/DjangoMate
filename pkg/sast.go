package pkg

import (
    "fmt"
    "os"
    "os/exec"

	"github.com/spf13/cobra"

)

var ScanCmd = &cobra.Command{
    Use:   "scan [path]",
    Short: "Scan Django project using Bandit for security vulnerabilities",
    Long: `This command runs Bandit to scan the entire Django project or a specific path
           for common security issues. If no path is provided, it scans the entire project.`,
    Args: cobra.MaximumNArgs(1),
    Run: func(cmd *cobra.Command, args []string) {

		var target string
        
		if len(args) > 0 {
            target = args[0]
        } else {
            cwd, err := os.Getwd()
            if err != nil {
                fmt.Println("Error getting current directory:", err)
                os.Exit(1)
            }
            target = cwd
        }

        fmt.Println("Running Bandit security scan on:", target)
        runBanditScan(target)
    },
}

func runBanditScan(target string) {

	cmd := exec.Command("cmd", "/C", "env\\Scripts\\bandit.exe" ,"-r", target)

	cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr

    err := cmd.Run()
    if err != nil {
        fmt.Println("Error running Bandit scan:", err)
        os.Exit(1)
    }
    fmt.Println("Bandit scan completed.")
}