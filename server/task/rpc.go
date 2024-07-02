package task

import (
	"context"
	"gsm/protobuf/taskpb"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

func Register(s *grpc.Server) {
	taskpb.RegisterTaskServiceServer(s, &server{})
}

type server struct {
	taskpb.UnsafeTaskServiceServer
}

func (*server) Create(ctx context.Context, req *taskpb.CreateRequest) (*taskpb.CreateResponse, error) {
	return Create(ctx, req)
}
func (*server) Get(ctx context.Context, req *taskpb.GetRequest) (*taskpb.GetResponse, error) {
	return Get(ctx, req)
}
func (*server) Update(ctx context.Context, req *taskpb.UpdateRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, Update(ctx, req)
}
func (*server) Delete(ctx context.Context, req *taskpb.DeleteRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, Delete(ctx, req)
}
func (*server) GetAll(ctx context.Context, req *taskpb.GetAllRequest) (*taskpb.GetAllResponse, error) {
	return GetAll(ctx, req)
}
