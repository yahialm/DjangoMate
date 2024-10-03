package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)



/*Run the developement server*/
var serverCmd = &cobra.Command{
    Use:   "server",
    Short: "Run the Django development server",
    Run: func(cmd *cobra.Command, args []string) {
        runserver(getProjectNameFromConfig())
    },
}

func runserver(projectName string) {
    // Get the current working directory (where manage.py should be located)
    projectDir, err := os.Getwd()
    if err != nil {
        fmt.Println("Error getting current working directory:", err)
    }

    manageFilePath := projectDir + "\\" + projectName + "\\manage.py"

    cmd := exec.Command("cmd", "/C", ".\\env\\Scripts\\python.exe", manageFilePath, "runserver")
    
    // Set the working directory where manage.py is located
    cmd.Dir = projectDir

    // Attach the process to the std output and error
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr

    // Run the command
    if err := cmd.Run(); err != nil {
        fmt.Println("Error running server:", err)
    }
}




/*Applies migrations to the database*/
var migCmd = &cobra.Command{
    Use:   "mig",
    Short: "Run python manage.py migrate",
    Run: func(cmd *cobra.Command, args []string) { 
        migrate(getProjectNameFromConfig())
    },
}

func migrate(projectName string) {
    projectDir, err := os.Getwd()
    if err != nil {
        fmt.Println("Error getting current working directory:", err)
    }

    manageFilePath := projectName + "\\manage.py"

    cmd := exec.Command("cmd", "/C", ".\\env\\Scripts\\python.exe", manageFilePath, "migrate")
    cmd.Dir = projectDir

    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr

    if err := cmd.Run(); err != nil {
        fmt.Println("Error running migrate:", err)
    }
}


/*Creates new migrations based on the changes detected in models*/
var mkmigCmd = &cobra.Command{
    Use:   "mkmig",
    Short: "Run python manage.py makemigrations",
    Run: func(cmd *cobra.Command, args []string) {
        makemigrations(getProjectNameFromConfig())
    },
}

func makemigrations(projectName string) {
    projectDir, err := os.Getwd()
    if err != nil {
        fmt.Println("Error getting current working directory:", err)
    }

    manageFilePath := projectName + "\\manage.py"

    cmd := exec.Command("cmd", "/C", ".\\env\\Scripts\\python.exe", manageFilePath, "makemigrations")
    cmd.Dir = projectDir

    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr

    if err := cmd.Run(); err != nil {
        fmt.Println("Error running makemigrations:", err)
    }
}


/*Lists all migrations and their status*/
var shmigCmd = &cobra.Command{
    Use:   "shmig",
    Short: "Run python manage.py showmigrations",
    Run: func(cmd *cobra.Command, args []string) {
        showmigrations(getProjectNameFromConfig())
    },
}

func showmigrations(projectName string) {
    projectDir, err := os.Getwd()
    if err != nil {
        fmt.Println("Error getting current working directory:", err)
    }

    manageFilePath := projectName + "\\manage.py"

    cmd := exec.Command("cmd", "/C", ".\\env\\Scripts\\python.exe", manageFilePath, "showmigrations")
    cmd.Dir = projectDir

    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr

    if err := cmd.Run(); err != nil {
        fmt.Println("Error running showmigrations:", err)
    }
}



/*Opens the database shell for direct database access*/
var dbshCmd = &cobra.Command{
    Use:   "dbsh",
    Short: "Run python manage.py dbshell",
    Run: func(cmd *cobra.Command, args []string) {
        dbshell(getProjectNameFromConfig())
    },
}

func dbshell(projectName string) {
    projectDir, err := os.Getwd()
    if err != nil {
        fmt.Println("Error getting current working directory:", err)
    }

    manageFilePath := projectName + "\\manage.py"

    cmd := exec.Command("cmd", "/C", ".\\env\\Scripts\\python.exe", manageFilePath, "dbshell")
    cmd.Dir = projectDir

    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr

    if err := cmd.Run(); err != nil {
        fmt.Println("Error running dbshell:", err)
    }
}


/*Opens the interactive Python shell with Django context*/
var shCmd = &cobra.Command{
    Use:   "sh",
    Short: "Run python manage.py shell",
    Run: func(cmd *cobra.Command, args []string) {
        shell(getProjectNameFromConfig())
    },
}

func shell(projectName string) {
    projectDir, err := os.Getwd()
    if err != nil {
        fmt.Println("Error getting current working directory:", err)
    }

    manageFilePath := projectName + "\\manage.py"

    cmd := exec.Command("cmd", "/C", ".\\env\\Scripts\\python.exe", manageFilePath, "shell")
    cmd.Dir = projectDir

    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr

    if err := cmd.Run(); err != nil {
        fmt.Println("Error running shell:", err)
    }
}

/*Run tests*/
var testCmd = &cobra.Command{
    Use:   "test",
    Short: "Run python manage.py test",
    Run: func(cmd *cobra.Command, args []string) {
        test(getProjectNameFromConfig())
    },
}

func test(projectName string) {
    projectDir, err := os.Getwd()
    if err != nil {
        fmt.Println("Error getting current working directory:", err)
    }

    manageFilePath := projectName + "\\manage.py"

    cmd := exec.Command("cmd", "/C", ".\\env\\Scripts\\python.exe", manageFilePath, "test")
    cmd.Dir = projectDir

    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr

    if err := cmd.Run(); err != nil {
        fmt.Println("Error running test:", err)
    }
}


/*Create super*/
var suCmd = &cobra.Command{
    Use:   "su",
    Short: "Run python manage.py createsuperuser",
    Run: func(cmd *cobra.Command, args []string) {
        createsu(getProjectNameFromConfig())
    },
}

func createsu(projectName string) {
    projectDir, err := os.Getwd()
    if err != nil {
        fmt.Println("Error getting current working directory:", err)
    }

    manageFilePath := projectName + "\\manage.py"

    cmd := exec.Command("cmd", "/C", ".\\env\\Scripts\\python.exe", manageFilePath, "createsuperuser")
    cmd.Dir = projectDir

    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr

    if err := cmd.Run(); err != nil {
        fmt.Println("Error running createsuperuser:", err)
    }
}