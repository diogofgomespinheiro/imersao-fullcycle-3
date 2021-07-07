package server

import (
	"log"
	"net"

	"github.com/diogofgomespinheiro/codebank/infrastructure/grpc/pb"
	"github.com/diogofgomespinheiro/codebank/infrastructure/grpc/service"
	"github.com/diogofgomespinheiro/codebank/usecases"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type GRPCServer struct {
	ProcessTransactionUseCase usecases.UseCaseTransaction
}

func NewGRPCServer() GRPCServer {
	return GRPCServer{}
}

func (g GRPCServer) Serve() {
	lis, err := net.Listen("tcp", "0.0.0.0:50052")
	if err != nil {
		log.Fatal("could not listen tcp port")
	}

	transactionService := service.NewTransactionService()
	transactionService.ProcessTransactionUseCase = g.ProcessTransactionUseCase
	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)
	pb.RegisterPaymentServiceServer(grpcServer, transactionService)

	grpcServer.Serve(lis)
}
