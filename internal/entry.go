package internal

import (
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

type Status int

const (
	None Status = iota
	Partial
	Complete
)
