package cmd

import (
    "github.com/spf13/cobra"
	"fmt"
	"os"
)

var rootCmd = &cobra.Command{
    Use:   "task-manager",
    Short: "A simple CLI Task Manager",
    Long:  "A simple CLI Task Manager to manage your tasks efficiently.",
}

func Execute() {
    if err := rootCmd.Execute(); err != nil {
        fmt.Fprintln(os.Stderr, err)
        os.Exit(1)
    }
}
