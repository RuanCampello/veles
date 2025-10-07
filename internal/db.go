package internal

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/sqlitedialect"
	"github.com/uptrace/bun/driver/sqliteshim"
)

type Db struct {
	*bun.DB
	ctx context.Context
}

type Habit struct {
	bun.BaseModel `bun:"table:habits"`

	Id   int64  `bun:"id, pk, autoincrement"`
	Name string `bun:"name, notnull, unique"`
}

func New(ctx context.Context) (*Db, error) {
	sql, err := sql.Open(sqliteshim.ShimName, "veles.db")
	if err != nil {
		return nil, fmt.Errorf("Failed to open database: %w", err)
	}

	db := bun.NewDB(sql, sqlitedialect.New())
	if err = createTables(ctx, db); err != nil {
		return nil, err
	}

	log.Println("database initialised successfully")
	return &Db{db, ctx}, nil
}

func createTables(ctx context.Context, db *bun.DB) error {
	panic("createTables not implemented")
}
