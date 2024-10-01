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
        runserver()
    },
}

func runserver() {
    // Execute: python manage.py runserver
    cmd := exec.Command("cmd", "/C", "env\\Scripts\\python.exe manage.py runserver")

    // Attach the process to the std output and error
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    if err := cmd.Run(); err != nil {
        fmt.Println("--------------------------------------------")
    }
}



/*Applies migrations to the database*/
var migCmd = &cobra.Command{
    Use:   "mig",
    Short: "Run python manage.py migrate",
    Run: func(cmd *cobra.Command, args []string) { 
        migrate()
    },
}

func migrate() {
    // Execute: python manage.py runserver
    cmd := exec.Command("cmd", "/C", "env\\Scripts\\python.exe manage.py migrate")

    // Attach the process to the std output and error
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    if err := cmd.Run(); err != nil {
        fmt.Println("--------------------------------------------")
    }
}


/*Creates new migrations based on the changes detected in models*/
var mkmigCmd = &cobra.Command{
    Use:   "mkmig",
    Short: "Run python manage.py makemigrations",
    Run: func(cmd *cobra.Command, args []string) {
        makemigrations()
    },
}

func makemigrations() {
    cmd := exec.Command("cmd", "/C", "env\\Scripts\\python.exe manage.py makemigrations")

    // Attach the process to the std output and error
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    if err := cmd.Run(); err != nil {
        fmt.Println("--------------------------------------------")
    }
}


/*Lists all migrations and their status*/
var shmigCmd = &cobra.Command{
    Use:   "shmig",
    Short: "Run python manage.py showmigrations",
    Run: func(cmd *cobra.Command, args []string) {
        showmigrations()
    },
}

func showmigrations() {
    cmd := exec.Command("cmd", "/C", "env\\Scripts\\python.exe manage.py showmigrations")

    // Attach the process to the std output and error
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    if err := cmd.Run(); err != nil {
        fmt.Println("--------------------------------------------")
    }
}



/*Opens the database shell for direct database access*/
var dbshCmd = &cobra.Command{
    Use:   "dbsh",
    Short: "Run python manage.py dbshell",
    Run: func(cmd *cobra.Command, args []string) {
        dbshell()
    },
}

func dbshell() {
    cmd := exec.Command("cmd", "/C", "env\\Scripts\\python.exe manage.py dbshell")

    // Attach the process to the std output and error
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    if err := cmd.Run(); err != nil {
        fmt.Println("--------------------------------------------")
    }
}


/*Opens the interactive Python shell with Django context*/
var shCmd = &cobra.Command{
    Use:   "sh",
    Short: "Run python manage.py shell",
    Run: func(cmd *cobra.Command, args []string) {
        shell()
    },
}

func shell() {
    cmd := exec.Command("cmd", "/C", "env\\Scripts\\python.exe manage.py shell")

    // Attach the process to the std output and error
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    if err := cmd.Run(); err != nil {
        fmt.Println("--------------------------------------------")
    }
}

/*Run tests*/
var testCmd = &cobra.Command{
    Use:   "test",
    Short: "Run python manage.py test",
    Run: func(cmd *cobra.Command, args []string) {
        test()
    },
}

func test() {
    cmd := exec.Command("cmd", "/C", "env\\Scripts\\python.exe manage.py test")

    // Attach the process to the std output and error
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    if err := cmd.Run(); err != nil {
        fmt.Println("--------------------------------------------")
    }
}


/*Create super*/
var suCmd = &cobra.Command{
    Use:   "su",
    Short: "Run python manage.py createsuperuser",
    Run: func(cmd *cobra.Command, args []string) {
        createsu()
    },
}

func createsu() {
    cmd := exec.Command("cmd", "/C", "env\\Scripts\\python.exe manage.py createsuperuser")

    // Attach the process to the std output and error
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    if err := cmd.Run(); err != nil {
        fmt.Println("--------------------------------------------")
    }
}