package grpc

import (
	"context"
	"log"
	"net"

	"github.com/bajalnyt/go-grpc-services-course/internal/rocket"
	rkt "github.com/bajalnyt/tutorial-protos/rocket"
	"google.golang.org/grpc"
)

// Handler will handle incoming gRPC requests
type Handler struct {
	RocketService RocketService
}

type RocketService interface {
	GetRocketById(ctx context.Context, id string) (rocket.Rocket, error)
	InsertRocket(ctx context.Context, rocket rocket.Rocket) (rocket.Rocket, error)
	DeleteRocketById(ctx context.Context, id string) error
}

func New(rktService RocketService) Handler {
	return Handler{
		RocketService: rktService,
	}
}

func (h Handler) Serve() error {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Println("could not start server")
		return err
	}

	grpcServer := grpc.NewServer()
	rkt.RegisterRocketServiceServer(grpcServer, &h)

	if err := grpcServer.Serve(lis); err != nil {
		log.Printf("failed to serve: %s\n", err)
		return err
	}

	return nil
}

func (h Handler) GetRocket(ctx context.Context, req *rkt.GetRocketRequest) (*rkt.GetRocketResponse, error) {
	return nil, nil
}
func (h Handler) AddRocket(ctx context.Context, req *rkt.AddRocketRequest) (*rkt.AddRocketResponse, error) {
	return nil, nil
}
func (h Handler) DeleteRocket(ctx context.Context, req *rkt.DeleteRocketRequest) (*rkt.DeleteRocketResponse, error) {
	return nil, nil
}
