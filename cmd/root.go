package cmd

import (
	"fmt"
	"os"

	"github.com/ruancampello/veles/cmd/habit"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "veles",
	Short: "An habit tracker for the neoclassicists",
}

func Exec() {
	rootCmd.AddCommand(habit.HabitCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
