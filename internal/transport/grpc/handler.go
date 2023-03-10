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
	log.Println("GetRocket gRPC endpoint hit")

	rocket, err := h.RocketService.GetRocketById(ctx, req.Id)
	if err != nil {
		log.Println("unable to fetch rocket by id")
		return &rkt.GetRocketResponse{}, err
	}

	return &rkt.GetRocketResponse{
		Rocket: &rkt.Rocket{
			Id:   rocket.ID,
			Name: rocket.Name,
			Type: rocket.Type,
		},
	}, nil
}

// AddRocket - adds a rocket to the database
func (h Handler) AddRocket(ctx context.Context, req *rkt.AddRocketRequest) (*rkt.AddRocketResponse, error) {
	log.Print("Add Rocket gRPC endpoint hit")
	newRkt, err := h.RocketService.InsertRocket(ctx, rocket.Rocket{
		ID:   req.Rocket.Id,
		Type: req.Rocket.Type,
		Name: req.Rocket.Name,
	})
	if err != nil {
		log.Print("failed to insert rocket into database")
		return &rkt.AddRocketResponse{}, err
	}
	return &rkt.AddRocketResponse{
		Rocket: &rkt.Rocket{
			Id:   newRkt.ID,
			Type: newRkt.Type,
			Name: newRkt.Name,
		},
	}, nil
}

// DeleteRocket - handler for deleting a rocket
func (h Handler) DeleteRocket(ctx context.Context, req *rkt.DeleteRocketRequest) (*rkt.DeleteRocketResponse, error) {
	log.Print("delete rocket gRPC endpoint hit")
	err := h.RocketService.DeleteRocketById(ctx, req.Rocket.Id)
	if err != nil {
		return &rkt.DeleteRocketResponse{}, err
	}
	return &rkt.DeleteRocketResponse{
		Status: "successfully delete rocket",
	}, nil
}
