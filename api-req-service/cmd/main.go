package main

import (
	"github.com/venomuz/task-iman-uz/api-req-service/config"
	pb "github.com/venomuz/task-iman-uz/api-req-service/genproto"
	"github.com/venomuz/task-iman-uz/api-req-service/pkg/db"
	"github.com/venomuz/task-iman-uz/api-req-service/pkg/logger"
	"github.com/venomuz/task-iman-uz/api-req-service/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
)

func main() {
	cfg := config.Load()

	log := logger.New(cfg.LogLevel, "req-service")
	defer logger.Cleanup(log)

	log.Info("main: sqlxConfig",
		logger.String("host", cfg.PostgresHost),
		logger.Int("port", cfg.PostgresPort),
		logger.String("database", cfg.PostgresDatabase))

	connDB, err := db.ConnectToDB(cfg)
	if err != nil {
		log.Fatal("sqlx connection to postgres error", logger.Error(err))
	}

	reqService := service.NewReqService(connDB, log)

	lis, err := net.Listen("tcp", cfg.RPCPort)
	if err != nil {
		log.Fatal("Error while listening: %v", logger.Error(err))
	}

	s := grpc.NewServer()
	pb.RegisterUserServiceServer(s, reqService)
	log.Info("main: server running",
		logger.String("port", cfg.RPCPort))
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatal("Error while listening: %v", logger.Error(err))
	}
}
