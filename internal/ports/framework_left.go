package ports

import (
	"appstruct/internal/adapters/framework/left/grpc/pb"
	"context"
)

type GRPCPort interface {
	Run()
	GetAdd(ctx context.Context, req *pb.OperationParameters) (*pb.Answer, error)
	GetSub(ctx context.Context, req *pb.OperationParameters) (*pb.Answer, error)
	GetMulti(ctx context.Context, req *pb.OperationParameters) (*pb.Answer, error)
	GetDiv(ctx context.Context, req *pb.OperationParameters) (*pb.Answer, error)
}
