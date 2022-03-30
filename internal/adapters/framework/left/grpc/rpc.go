package grpc

import (
	"appstruct/internal/adapters/framework/left/grpc/pb"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func calculate(ctx context.Context, parameters *pb.OperationParameters, calc func(int32, int32) (int32, error)) (*pb.Answer, error) {
	if parameters.GetA() == 0 || parameters.GetB() == 0 {
		return nil, status.Error(codes.InvalidArgument, "missing required parameters")
	}

	answer, err := calc(parameters.A, parameters.B)
	if err != nil {
		return nil, status.Error(codes.Internal, "unexpected error")
	}
	return &pb.Answer{Value: answer}, nil
}

func (a Adapter) GetAdd(ctx context.Context, parameters *pb.OperationParameters) (*pb.Answer, error) {
	return calculate(ctx, parameters, a.api.GetAdd)
}

func (a Adapter) GetSub(ctx context.Context, parameters *pb.OperationParameters) (*pb.Answer, error) {
	return calculate(ctx, parameters, a.api.GetSub)
}

func (a Adapter) GetMulti(ctx context.Context, parameters *pb.OperationParameters) (*pb.Answer, error) {
	return calculate(ctx, parameters, a.api.GetMulti)
}

func (a Adapter) GetDiv(ctx context.Context, parameters *pb.OperationParameters) (*pb.Answer, error) {
	return calculate(ctx, parameters, a.api.GetDiv)
}
