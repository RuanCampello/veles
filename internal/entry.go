package internal

import (
	"time"

	"github.com/uptrace/bun"
)

type Entry struct {
	bun.BaseModel `bun:"table:entry"`

	id        int64     `bun:"id, pk, autoincrement"`
	habitId   int64     `bun:"habit_id, notnull"`
	date      string    `bun:"date, notnull"`
	status    Status    `bun:"status, notnull"`
	note      string    `bun:"note"`
	createdAt time.Time `bun:"created_at, default:current_timestamp"`
	habit     *Habit    `bun:"rel:belongs-to,join:habit_id=id"`
}

type Status int

const (
	None Status = iota
	Partial
	Complete
)
