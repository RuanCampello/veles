package habit

import (
	"github.com/spf13/cobra"
)

var HabitCmd = &cobra.Command{
	Use:   "habit",
	Short: "Manage your habits",
}

func init() {
	HabitCmd.AddCommand(AddCmd)
	HabitCmd.AddCommand(ListCmd)
}
