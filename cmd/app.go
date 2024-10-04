package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)


var appCmd = &cobra.Command{
    Use:   "app [app name]",
    Short: "Create new application and register the created app to the settings.py configuration",
	Args:  cobra.MinimumNArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
		createApp(args[0], getProjectNameFromConfig())
    },
}

func createApp(appName, projectName string) {
    fmt.Printf("Creating app: %s...\n", appName)
	cmd := exec.Command("..\\env\\Scripts\\python.exe", "manage.py", "startapp", appName)
    cmd.Dir = fmt.Sprintf(".\\%s", projectName)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
    if err := cmd.Run(); err != nil {
        fmt.Println("Error:", err)
        return
    }

    settingsPath := fmt.Sprintf("./%s/%s/settings.py", projectName, projectName)
    addAppToSettings(settingsPath, appName)
}

func addAppToSettings(settingsPath, appName string) {
    
    fileContent, err := os.ReadFile(settingsPath)
    if err != nil {
        fmt.Println("Error reading settings.py:", err)
        return
    }

    fileString := string(fileContent)

	// Find the INSTALLED_APPS variable
    startIndex := strings.Index(fileString, "INSTALLED_APPS = [")
    if startIndex == -1 {
        fmt.Println("INSTALLED_APPS not found in settings.py")
        return
    }

	// Get the index of the closing bracket
    endIndex := strings.Index(fileString[startIndex:], "]")
    if endIndex == -1 {
        fmt.Println("Error finding end of INSTALLED_APPS")
        return
    }
    endIndex += startIndex  // Adjust the index for the whole file

	// AppName is appName but the first letter is uppercase
	var AppName string = strings.ToUpper(appName[:1]) + appName[1:]

    newAppEntry := fmt.Sprintf("    '%s.apps.%sConfig',\n", appName, AppName)
    updatedFile := fileString[:endIndex] + newAppEntry + fileString[endIndex:]

    // Step 5: Write the updated content back to settings.py
    err = os.WriteFile(settingsPath, []byte(updatedFile), 0644)
    if err != nil {
        fmt.Println("Error writing to settings.py:", err)
        return
    }

    fmt.Printf("App '%s' has been added to INSTALLED_APPS in settings.py.\n", appName)
}