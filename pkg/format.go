package pkg

import (

	"github.com/spf13/cobra"

	"os"
	"os/exec"
	"fmt"

)

var FormatCmd = &cobra.Command{
    Use:   "format [path]",
    Short: "Formats Python code using Black",
    Long: `This command formats your Django project using Black, ensuring consistent code formatting.`,
    Args: cobra.MaximumNArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
        var target string
        if len(args) > 0 {
            target = args[0]
        } else {
            cwd, _ := os.Getwd()
            target = cwd
        }
        formatCode(target)
    },
}


func formatCode(target string) {

	cmd := exec.Command("cmd", "/C", "env\\Scripts\\black.exe" , target)

	cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr

    // Run the command
    err := cmd.Run()
    if err != nil {
        fmt.Println("Error running black formatter:", err)
        os.Exit(1)
    }
    fmt.Println("Format completed.")
}