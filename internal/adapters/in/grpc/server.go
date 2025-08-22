package grpc

import (
	"ai-agent-manager/api/proto/agentmanager"
	"ai-agent-manager/internal/platform/config"

	"github.com/jackc/pgx/v5/pgxpool"
	logx "github.com/junhekdevsru/ai-common-lib/logx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/emptypb"
)

type healthImpl struct {
	agent_manager.UnimplementedHealthServiceServer
}

func (h *healthImpl) Ping(*emptypb.Empty) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}

func NewServer(log logx.Logger, _ config.Config, _ *pgxpool.Pool) *grpc.Server {
	s := grpc.NewServer(
		grpc.UnaryInterceptor(logx.GrpcServerInterceptor(log)),
	)
	reflection.Register(s)
	agent_manager.RegisterHealthServiceServer(s, &healthImpl{})
	return s
}
