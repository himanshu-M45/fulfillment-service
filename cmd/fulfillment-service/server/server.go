package server

import (
	"errors"
	"fulfillment-service/internal/app/fs/controllers"
	"fulfillment-service/internal/app/fs/db"
	proto "fulfillment-service/proto"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/golang-migrate/migrate/v4/source/github"
	_ "github.com/lib/pq"
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

	// Run database migrations
	runMigrations(dataSourceName)

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

func runMigrations(dataSourceName string) {
	database := db.GetDB()
	if database == nil {
		log.Fatalf("Failed to connect to the database")
	}

	driver, err := postgres.WithInstance(database, &postgres.Config{})
	if err != nil {
		log.Fatalf("Failed to create migration driver: %v", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file:///Users/himanshu.m_ftc/Desktop/GoProjects/fulfillment-service/internal/app/fs/db/migrations",
		"postgres", driver)
	if err != nil {
		log.Fatalf("Failed to create migrate instance: %v", err)
	}

	if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	log.Println("Database migrations applied successfully")
}
