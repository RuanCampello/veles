package internal

import (
	"context"
	"fmt"
	"time"

	"github.com/uptrace/bun"
)

type Entry struct {
	bun.BaseModel `bun:"table:entry"`

	Id        int64     `bun:"id,pk,autoincrement"`
	HabitId   int64     `bun:"habit_id,notnull"`
	Date      string    `bun:"date,notnull"`
	Status    Status    `bun:"status,notnull"`
	Note      string    `bun:"note"`
	CreatedAt time.Time `bun:"created_at,default:current_timestamp"`
	Habit     *Habit    `bun:"rel:belongs-to,join:habit_id=id"`
}

type Status string

const (
	None     Status = "none"
	Partial  Status = "partial"
	Complete Status = "complete"
)

func (db *Db) CreateEntry(ctx context.Context, habitId int64, progress Status, comment string) (*Entry, error) {
	entry := &Entry{
		HabitId: habitId,
		Status:  progress,
		Note:    comment,
	}
	// this should error cause we don't pass the date

	if _, err := db.NewInsert().Model(entry).Exec(ctx); err != nil {
		return nil, fmt.Errorf("failed to create new entry: %w", err)
	}

	return entry, nil
}
