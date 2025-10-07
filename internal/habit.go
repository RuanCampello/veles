package internal

import (
	"context"
	"time"

	"github.com/uptrace/bun"
)

type Habit struct {
	bun.BaseModel `bun:"table:habit"`

	id          int64     `bun:"id, pk, autoincrement"`
	name        string    `bun:"name, notnull, unique"`
	description string    `bun:"description"`
	createdAt   time.Time `bun:"created_at, default:current_timestamp"`
	entries     []*Entry  `bun:"rel:has-many, join:id=habit_id"`
}

func CreateHabit(ctx context.Context, db *Db, name, description string) (*Habit, error) {
	// habit := &Habit{
	// 	name:        name,
	// 	description: description,
	// }

	panic("CreateHabit not yet implemented")
}
