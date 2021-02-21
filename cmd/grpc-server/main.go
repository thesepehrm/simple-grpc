package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net"
	"time"

	"github.com/thesepehrm/simple-grpc/common"
	"github.com/thesepehrm/simple-grpc/pb/update"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

const (
	port = ":3000"
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	users *common.UpdateUsers
	update.UnimplementedUpdateServer
}

func (s *server) Login(ctx context.Context, req *update.LoginRequest) (*update.LoginResponse, error) {
	token := s.users.Login(req.GetUsername(), req.GetPassword())
	return &update.LoginResponse{
		Token: token,
	}, nil
}
func (s *server) Logout(ctx context.Context, req *update.LogoutRequest) (*emptypb.Empty, error) {
	return new(emptypb.Empty), s.users.Logout(req.Token)
}

func (s *server) ServerPromotions(_ *emptypb.Empty, stream update.Update_ServerPromotionsServer) error {
	for range time.Tick(time.Second) {
		err := stream.Send(&update.UpdateStreamResponse{
			Timestamp: timestamppb.Now(),
			Update: &update.UpdateStreamResponse_Update{
				Type: "promotion",
				Text: fmt.Sprintf("%d%%", rand.Intn(100)),
			},
		})
		if err != nil {
			log.Print(err)
		}

	}
	return nil
}

func newServer() *server {
	return &server{
		users: common.NewUpdateUsers(),
	}
}

func main() {

	lis, err := net.Listen("tcp", fmt.Sprintf("localhost"+port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption

	grpcServer := grpc.NewServer(opts...)
	update.RegisterUpdateServer(grpcServer, newServer())
	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatalf(err.Error())
	}

}
