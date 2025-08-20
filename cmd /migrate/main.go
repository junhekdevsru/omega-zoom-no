package main

import (
	"ai-agent-manager/internal/platform/config"
	"ai-agent-manager/internal/platform/db"
	"context"
	"flag"
	"log"
)

func main() {
	direction := flag.String("direction", "up", "up|down")
	flag.Parse()

	cfg := config.MustLoad()
	ctx := context.Background()

	pool, err := db.Connect(ctx, cfg)
	if err != nil {
		log.Fatalf("db connect: %v", err)
	}
	defer pool.Close()

	if err := db.ApplyMigrations(ctx, pool, cfg.Migrations, *direction); err != nil {
		log.Fatalf("migrate %s: %v", *direction, err)
	}
	log.Printf("migrations %s done", *direction)
}
