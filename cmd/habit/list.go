package habit

import (
	"context"
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/ruancampello/veles/internal"
	"github.com/spf13/cobra"
)

var ListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all habits",
	Run: func(cmd *cobra.Command, args []string) {

		ctx := context.Background()
		db := internal.NewDb(ctx)
		defer db.Close()

		habits, err := db.ListHabits(ctx)
		if err != nil {
			fmt.Printf("Failed to get habits: %v", err)
			return
		}

		if len(habits) == 0 {
			fmt.Println("No habit found. Try to add a new one :D.")
		}

		w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
		fmt.Fprintln(w, "ID\tName\tDescription\tCreated At")

		for _, habit := range habits {
			created := habit.CreatedAt.Format("2006-01-02")
			fmt.Fprintf(w, "%d\t%s\t%s\t%s\n", habit.Id, habit.Name, habit.Description, created)
		}

		w.Flush()
	},
}
