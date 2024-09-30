package cmd


import (
    "os/exec"
    "github.com/spf13/cobra"
    "fmt"
)


var appCmd = &cobra.Command{
    Use:   "app",
    Short: "Create new application and register the created app to the settings.py configuration",
    Run: func(cmd *cobra.Command, args []string) {
		createApp()
    },
}

func createApp() {
    // Execute: python manage.py runserver
    out, err := exec.Command("python", "manage.py", "startapp").Output()
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    fmt.Println(string(out))
}