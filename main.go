package main

import (
	"github.com/spf13/cobra"
	"log"
	"os"
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "task-manager",
		Short: "A simple CLI Task Manager",
		Long:  "A simple CLI Task Manager to manage your tasks efficiently.",
	}

	if err := rootCmd.Execute(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}