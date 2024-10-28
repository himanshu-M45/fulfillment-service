package server

import (
	"fulfillment-service/internal/app/fs/controllers"
	"fulfillment-service/internal/app/fs/db"
	proto "fulfillment-service/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

const (
	port = ":50052"
)

func RunServer() {
	// Database connection
	dataSourceName := "user=fulfill_user password=your_password dbname=fulfill_db sslmode=disable"
	db.InitDB(dataSourceName)
	if db.GetDB() == nil {
		log.Fatalf("Failed to connect to the database")
	}

	// gRPC server setup
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	proto.RegisterFulfillmentServiceServer(grpcServer, &controllers.Server{})

	log.Printf("Server listening at %v", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
