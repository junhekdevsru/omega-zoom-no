package main

import (
	"context"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	logx "github.com/junhekdevsru/ai-common-lib/logx"

	"ai-agent-manager/internal/adapters/in/grpc"
	"ai-agent-manager/internal/platform/config"
	"ai-agent-manager/internal/platform/db"
	"ai-agent-manager/internal/platform/observability"

	grpc_health "google.golang.org/grpc/health"
	grpc_health_v1 "google.golang.org/grpc/health/grpc_health_v1"
)

func main() {
	base := observability.NewLogger() // zerolog.Logger
	l := logx.New(base)               // logx.Logger

	cfg := config.MustLoad()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	pool, err := db.Connect(ctx, cfg)
	if err != nil {
		base.Fatal().Err(err).Msg("db connect failed")
	}
	defer pool.Close()

	srv := grpc.NewServer(l, cfg, pool)

	health := grpc_health.NewServer()
	grpc_health_v1.RegisterHealthServer(srv, health)

	lis, err := net.Listen("tcp", cfg.GRPCAddr)
	if err != nil {
		base.Fatal().Err(err).Msg("listen failed")
	}

	go func() {
		l.Info("gRPC listening", "addr", cfg.GRPCAddr)
		if err := srv.Serve(lis); err != nil {
			l.Error(err, "gRPC server stopped")
			cancel()
		}
	}()

	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc, os.Interrupt, syscall.SIGTERM)
	select {
	case <-sigc:
		l.Info("shutdown signal")
	case <-ctx.Done():
	}

	stCtx, stCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer stCancel()

	srv.GracefulStop()
	if stCtx.Err() != nil {
		srv.Stop()
	}
	l.Info("bye")
}
