package habit

import (
	"context"
	"fmt"

	"github.com/ruancampello/veles/internal"
	"github.com/spf13/cobra"
)

var AddCmd = &cobra.Command{
	Use:   "add [name]",
	Short: "Creates a new habit",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		description, _ := cmd.Flags().GetString("description")

		ctx := context.Background()
		db := internal.NewDb(ctx)
		defer db.Close()

		habit, err := db.CreateHabit(ctx, name, description)
		if err != nil {
			fmt.Printf("Failed to create habit: %v\n", err)
			return
		}

		fmt.Printf("Habit %s's created successfully\n", habit.Name)
	},
}

func init() {
	AddCmd.Flags().StringP("description", "d", "", "Habit description")
}
