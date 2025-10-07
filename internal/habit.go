package internal

import (
	"context"
	"fmt"
	"time"

	"github.com/uptrace/bun"
)

type Habit struct {
	bun.BaseModel `bun:"table:habit"`

	Id          int64     `bun:"id,pk,autoincrement"`
	Name        string    `bun:"name,notnull,unique"`
	Description string    `bun:"description"`
	CreatedAt   time.Time `bun:"created_at,default:current_timestamp"`
	Entries     []*Entry  `bun:"rel:has-many,join:id=habit_id"`
}

func CreateHabit(ctx context.Context, db *Db, name, description string) (*Habit, error) {
	habit := &Habit{
		Name:        name,
		Description: description,
	}

	if _, err := db.NewInsert().Model(&habit).Exec(ctx); err != nil {
		return nil, fmt.Errorf("Failed to create habit: %w", err)
	}

	return habit, nil
}
