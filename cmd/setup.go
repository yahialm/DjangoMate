package cmd

import (
    "bufio"
    "fmt"
    "os"
    "os/exec"
    "strings"

    "github.com/spf13/cobra"
)

var setupCmd = &cobra.Command{
    Use:   "setup",
    Short: "Setup your django project with one command. Choose your database and get your settings.py automatically configured",
    Run: func(cmd *cobra.Command, args []string) {
        setup()
    },
}

func init() {
    rootCmd.AddCommand(setupCmd)
}

// Setup the project environment
func setup() {
    projectName := getProjectName()
    createConfigFile(projectName)
    createVirtualEnv()
	activateVenvWindows()
    installDjango()
    startDjangoProject(projectName)
    dbConfig(projectName)
    installDjangoEnviron()
    dbEnvVars(projectName)
}

// Get the name of the project
func getProjectName() string {
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Enter the Django project name: ")
    projectName, _ := reader.ReadString('\n')
    return strings.TrimSpace(projectName)
}

// Creates a config file with a warning and project name
func createConfigFile(projectName string) {
    configContent := fmt.Sprintf(
        "Warning: Do not modify this file\n%s", projectName)

    err := os.WriteFile("config", []byte(configContent), 0644)
    if err != nil {
        fmt.Println("Error writing to config file:", err)
        return
    }

    fmt.Println("Config file created with project name.")
}

// Read the project name from the conf file
func getProjectNameFromConfig() string {
    fileContent, err := os.ReadFile("config")
    if err != nil {
        fmt.Println("Error reading the config file:", err)
        return ""
    }

    // Split the content into lines
    lines := strings.Split(string(fileContent), "\n")

    // The project name is on the second line
    if len(lines) > 1 {
        return strings.TrimSpace(lines[1])
    }

    fmt.Println("Project name not found in the config file.")
    return ""
}

// Create python virtual environment using 'virtualenv'
func createVirtualEnv() {
    if _, err := exec.LookPath("pip"); err != nil {
        fmt.Println("pip is not installed. Please install pip first.")
        return
    }

    if _, err := exec.LookPath("virtualenv"); err != nil {
        fmt.Println("virtualenv is not installed. Installing...")
        exec.Command("pip", "install", "virtualenv").Run()
    }

    fmt.Println("Creating virtual environment...")
    exec.Command("virtualenv", "env").Run()
    fmt.Println("Virtual environment created.")
}

func activateVenvWindows() {
    cmd := exec.Command("cmd", "/C", "env\\Scripts\\activate.bat")
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    err := cmd.Run()
    if err != nil {
        fmt.Println("Can't activate virtual environment")
        fmt.Println("More details: ", err)
    } else {
        fmt.Println("Virtual environment activated successfully.")
    }
}

// Install django using pip.exe from the virtual environment
func installDjango() {

    fmt.Println("Installing django...")

    cmd := exec.Command("cmd", "/C", "env\\Scripts\\pip.exe install django")

    // Attach the process to the std output and error
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    if err := cmd.Run(); err != nil {
        fmt.Println("Django installation failed")
    } else {
        fmt.Println("Django installation completed")
    }
}

func installDjangoEnviron() {

    fmt.Println("Installing django-environ...")

    cmd := exec.Command("cmd", "/C", "env\\Scripts\\pip.exe install django-environ")

    // Attach the process to the std output and error
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    if err := cmd.Run(); err != nil {
        fmt.Println("Django-environ installation failed")
    } else {
        fmt.Println("Django-environ installation completed")
    }

    
}

// Generate the template of the project
// After activating the virtualenv, we use "pip install django" to install django
// TODO: django-admin startproject <project_name> is executed in the global environment
func startDjangoProject(projectName string) {
    fmt.Println("Creating Django project...")
    cmd := exec.Command("cmd", "/C", "env\\Scripts\\django-admin.exe", "startproject", projectName)
    err := cmd.Run()
    if err != nil {
        fmt.Println("Error creating Django project:", err)
        return
    }
    fmt.Println("Django project created.")
}

// Configure database and handle virtual environment
func dbConfig(projectName string) {
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Enter the database you plan to use (sqlite3, mysql, postgresql): ")
    dbChoice, _ := reader.ReadString('\n')
    dbChoice = strings.TrimSpace(dbChoice)

    virtualEnvPath := "./env/bin/pip" // Change path to virtual env pip for Unix-based systems
    if _, err := os.Stat("env/Scripts/pip.exe"); err == nil {
        virtualEnvPath = "env/Scripts/pip.exe" // For Windows systems
    }

    if dbChoice != "sqlite3" {
        fmt.Printf("Installing database driver for %s inside virtual environment...\n", dbChoice)
        var pkg string
        if dbChoice == "mysql" {
            pkg = "mysqlclient"
        } else if dbChoice == "postgresql" {
            pkg = "psycopg2-binary"
        }
        if pkg != "" {
            // Install the package using pip from the virtual environment
            cmd := exec.Command(virtualEnvPath, "install", pkg)
            cmd.Stderr = os.Stderr
            cmd.Stdout = os.Stdout
            err := cmd.Run()
            if err != nil {
                fmt.Println("Error installing the package:", err)
                return
            }
        }
    }

    // Modify settings.py
    settingsPath := fmt.Sprintf("./%s/%s/settings.py", projectName, projectName)
    fileContent, err := os.ReadFile(settingsPath)
    if err != nil {
        fmt.Println("Error reading settings.py:", err)
        return
    }

    fileString := string(fileContent)

    // Remove existing DATABASES configuration if it exists
    startIndex := strings.Index(fileString, "DATABASES = {")
    if startIndex != -1 {
        // Use a brace counter to find the end of the block
        braceCount := 1
        endIndex := startIndex + len("DATABASES = {")
        for endIndex < len(fileString) && braceCount > 0 {
            if fileString[endIndex] == '{' {
                braceCount++
            } else if fileString[endIndex] == '}' {
                braceCount--
            }
            endIndex++
        }

        // Ensure we remove any trailing newline or extra spaces after the closing brace
        fileString = fileString[:startIndex] + fileString[endIndex:]
    }

    // Create the new database configuration
    dbConfig := fmt.Sprintf(`

# Import the os module to use environment variables 
import environ

env = environ.Env()
environ.Env.read_env()

# Database configuration
DATABASES = {
    'default': {
        'ENGINE': 'django.db.backends.%s',
        'NAME': env('%s_DB_NAME'),
        'USER': env('%s_DB_USER'),
        'PASSWORD': env('%s_DB_PASSWORD'),
        'HOST': env('%s_DB_HOST'),
        'PORT': env('%s_DB_PORT'),
    }
}
`, dbChoice, strings.ToUpper(projectName), strings.ToUpper(projectName), strings.ToUpper(projectName), strings.ToUpper(projectName), strings.ToUpper(projectName))

    // Append the new configuration
    fileString += dbConfig

    // Write the changes back to settings.py
    if err = os.WriteFile(settingsPath, []byte(fileString), 0644); err != nil {
        fmt.Println("Error writing to settings.py:", err)
        return
    }

    fmt.Println("Database configured in settings.py.")
}

// os.Setenv will not work properly cuz it set environment variables for the current process
// Once the current process/program exit the env vars will b immediatly deleted
// The usage of .env file give the possibility to persist the variables beyond program life
// dbEnvVars collects DB credentials and saves them to a .env file
func dbEnvVars(projectName string) {
    reader := bufio.NewReader(os.Stdin)

    // Get DB_NAME
    fmt.Print("Enter DB_NAME: ")
    dbName, _ := reader.ReadString('\n')
    dbName = strings.TrimSpace(dbName)

    // Get DB_USER
    fmt.Print("Enter DB_USER: ")
    dbUser, _ := reader.ReadString('\n')
    dbUser = strings.TrimSpace(dbUser)

    // Get DB_PASSWORD
    fmt.Print("Enter DB_PASSWORD: ")
    dbPassword, _ := reader.ReadString('\n')
    dbPassword = strings.TrimSpace(dbPassword)

    // Get DB_HOST
    fmt.Print("Enter DB_HOST: ")
    dbHost, _ := reader.ReadString('\n')
    dbHost = strings.TrimSpace(dbHost)

    // Get DB_PORT
    fmt.Print("Enter DB_PORT: ")
    dbPort, _ := reader.ReadString('\n')
    dbPort = strings.TrimSpace(dbPort)

    // Build the .env file content
    envContent := fmt.Sprintf(
        `%s_DB_NAME=%s
%s_DB_USER=%s
%s_DB_PASSWORD=%s
%s_DB_HOST=%s
%s_DB_PORT=%s`, strings.ToUpper(projectName), dbName, strings.ToUpper(projectName), dbUser, strings.ToUpper(projectName), dbPassword, strings.ToUpper(projectName), dbHost, strings.ToUpper(projectName), dbPort)

    // Create or overwrite the .env file
    envFilePath := ".env"
    err := os.WriteFile(envFilePath, []byte(envContent), 0644)
    if err != nil {
        fmt.Println("Error writing to .env file:", err)
        return
    }

    fmt.Println("Environment variables saved to .env file.")
}