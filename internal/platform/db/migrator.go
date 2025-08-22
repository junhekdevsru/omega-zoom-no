package db

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/jackc/pgx/v5/pgxpool"
)

func ApplyMigrations(ctx context.Context, pool *pgxpool.Pool, dir string, direction string) error {
	entries, err := filepath.Glob(filepath.Join(dir, "*.sql"))
	if err != nil {
		return err
	}
	sort.Strings(entries)

	tx, err := pool.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	if _, err := tx.Exec(ctx, `
CREATE TABLE IF NOT EXISTS schema_migrations(
  id SERIAL PRIMARY KEY,
  name TEXT NOT NULL UNIQUE,
  applied_at TIMESTAMPTZ NOT NULL DEFAULT now()
)`); err != nil {
		return err
	}

	switch strings.ToLower(direction) {
	case "up":
		for _, f := range entries {
			name := filepath.Base(f)
			var exists bool
			if err := tx.QueryRow(ctx, "SELECT EXISTS (SELECT 1 FROM schema_migrations WHERE name=$1)", name).Scan(&exists); err != nil {
				return err
			}
			if exists {
				continue
			}
			b, err := os.ReadFile(f)
			if err != nil {
				return err
			}
			if _, err := tx.Exec(ctx, string(b)); err != nil {
				return fmt.Errorf("migration %s failed: %w", name, err)
			}
			if _, err := tx.Exec(ctx, "INSERT INTO schema_migrations(name) VALUES($1)", name); err != nil {
				return err
			}
		}
	case "down":
		return fmt.Errorf("down not implemented in simple migrator")
	default:
		return fmt.Errorf("unknown direction: %s", direction)
	}

	return tx.Commit(ctx)
}
