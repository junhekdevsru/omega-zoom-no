package main

import (
	"context"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"ai-agent-manager/internal/adapters/in/grpc"
	"ai-agent-manager/internal/platform/config"
	"ai-agent-manager/internal/platform/db"
	"ai-agent-manager/internal/platform/observability"

	grpc_health "google.golang.org/grpc/health"
	grpc_health_v1 "google.golang.org/grpc/health/grpc_health_v1"
)

func main() {
	log := observability.NewLogger()
	cfg := config.MustLoad()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// DB connect
	pool, err := db.Connect(ctx, cfg)
	if err != nil {
		log.Fatal.Err(err).Msg("db connection failed")
	}
	defer pool.Close()

	// gRPC server
	srv := grpc.NewServer(log, cfg, pool)

	// standart grpc-health
	health := grpc.health.NewServer()
	grpc_health_v1.RegisterHealthServer(srv, health)

	lis, err := net.Listen("tcp", cfg.GRPCAddr)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to listen")
	}

	// graceful shutdown
	go func() {
		log.Info().Str("addr", cfg.GRPCAddr).Msg("gRPC listening")
		if err := srv.Serve(lis); err != nil {
			log.Error().Err(err).Msg("gRPC server stopped")
			cancel()
		}
	}()

	// handle signals
	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc, os.Interrupt, syscall.SIGTERM)
	select {
	case <-sigc:
		log.Info().Msg("shutdown signal")
	case <-ctx.Done():
	}

	stCtx, stCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer stCancel()

	srv.GracefulStop()
	if stCtx.Err() == context.DeadlineExceeded {
		srv.Stop()
	}
	log.Info().Msg("bye")
}
