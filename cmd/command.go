package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

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

var mkmigCmd = &cobra.Command{
    Use:   "mkmig",
    Short: "Run the Django development server",
    Run: func(cmd *cobra.Command, args []string) {
        makemigrations()
    },
}

func makemigrations() {
    cmd := exec.Command("cmd", "/C", "env\\Scripts\\python.exe manage.py migrate")

    // Attach the process to the std output and error
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    if err := cmd.Run(); err != nil {
        fmt.Println("--------------------------------------------")
    }
}


// var appCmd = &cobra.Command{
//     Use:   "app",
//     Short: "Run the Django development server",
//     Run: func(cmd *cobra.Command, args []string) {
//         // Execute: python manage.py runserver
//         out, err := exec.Command("python", "manage.py", "runserver").Output()
//         if err != nil {
//             fmt.Println("Error:", err)
//             return
//         }
//         fmt.Println(string(out))
//     },
// }