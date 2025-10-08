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

func NewDb(ctx context.Context) (*Db, error) {
	sql, err := sql.Open(sqliteshim.ShimName, "file:veles.db")
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
	if _, err := db.NewCreateTable().Model((*Habit)(nil)).IfNotExists().Exec(ctx); err != nil {
		return fmt.Errorf("Couldn't create habits table: %w", err)
	}

	if _, err := db.NewCreateTable().Model((*Entry)(nil)).IfNotExists().Exec(ctx); err != nil {
		return fmt.Errorf("Couldn't create entries table: %w", err)
	}

	return nil
}
