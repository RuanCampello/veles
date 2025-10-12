package entry

import (
	"context"
	"fmt"
	"strconv"

	"github.com/ruancampello/veles/internal"
	"github.com/spf13/cobra"
)

var AddCmd = &cobra.Command{
	Use:   "add [habit-name-or-id]",
	Short: "Creates a new entry to a habit",
	Run: func(cmd *cobra.Command, args []string) {
		id := args[0]

		progress, _ := cmd.Flags().GetString("progress")
		comment, _ := cmd.Flags().GetString("comment")

		progressEnum := internal.Status(progress)
		if progressEnum != internal.None &&
			progressEnum != internal.Partial &&
			progressEnum != internal.Complete {
			fmt.Printf("Error: progress status must be one of: none, partial, complete\n")
			return
		}

		ctx := context.Background()
		db := internal.NewDb(ctx)
		defer db.Close()

		var habit *internal.Habit
		if id, err := strconv.ParseInt(id, 10, 64); err != nil {
			habit, err := db.GetHabitById(ctx, id)
			if err != nil {
				fmt.Printf("Habit with ID %s was not found: %v", id, err)
				return
			}
		} else {
			habit, err := db.GetHabitById(ctx, id)
			if err != nil {
				fmt.Printf("Habit with name %s was not found: %v", id, err)
				return
			}
		}

		entry, err := db.CreateEntry(ctx, habit.Id, progressEnum, comment)
		if err != nil {
			fmt.Printf("Error: %v", err)
		}

		fmt.Printf("Entry created successfully for habit '%s' (ID: %d)\n", habit.Name, entry.Id)
	},
}
