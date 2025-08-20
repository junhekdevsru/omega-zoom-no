package grpc

import (
	"ai-agent-manager/api/proto/agentmgr"
	"ai-agent-manager/internal/platform/config"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/emptypb"
)

type healthImpl struct {
	agentmgr.UnimplementedHealthServiceServer
}

func (h *healthImpl) Ping(*emptypb.Empty) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}

func NewServer(log zerolog.Logger, cfg config.Config, pool *pgxpool.Pool) *grpc.Server {
	s := grpc.NewServer(
	// тут позже добавим интерсепторы: logging, tracing, recovery, auth
	)

	// dev‑reflection (удобно для grpcurl)
	reflection.Register(s)

	// register minimal health service (наш)
	agentmgr.RegisterHealthServiceServer(s, &healthImpl{})

	return s
}
