package main

import (
	"log"

	"github.com/bajalnyt/go-grpc-services-course/internal/db"
	"github.com/bajalnyt/go-grpc-services-course/internal/rocket"
	"github.com/bajalnyt/go-grpc-services-course/internal/transport/grpc"
)

// Run will init and start the gRPC server
func Run() error {
	rocketStore, err := db.New()
	if err != nil {
		return err
	}

	err = rocketStore.MigrateDB()
	if err != nil {
		log.Println(err.Error())
		return err
	}

	rktService := rocket.New(rocketStore)

	rktHandler := grpc.New(rktService)

	if err := rktHandler.Serve(); err != nil {
		log.Println("failed to start server")
		return err
	}

	return nil
}

func main() {
	if err := Run(); err != nil {
		log.Fatal(err)
	}
}
