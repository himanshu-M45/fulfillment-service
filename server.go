package main

import (
	"fulfillment-service/db"
	proto "fulfillment-service/proto"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"log"
	"net"
)

const (
	port = ":50052"
)

func main() {
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
	proto.RegisterFulfillmentServiceServer(grpcServer, &Server{})

	log.Printf("Server listening at %v", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
