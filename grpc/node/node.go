package grpcnode

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"

	"demo-transaction/config"
	transactionpb "demo-transaction/proto/models/transaction"
)

// Node ...
type Node struct{}

// GetTransactionDetailByUserID ...
func (s *Node) GetTransactionDetailByUserID(ctx context.Context, req *transactionpb.GetTransactionDetailByUserIDRequest) (*transactionpb.GetTransactionDetailByUserIDResponse, error) {
	var(
		userID = req.GetUserID()
	)

	// Get user by id
	result, err := getTransactionDetailByUserID(userID)

	return result, err
}

// Start ...
func Start() {
	envVars := config.GetEnv()
	transactionPort := config.GetEnv().GRPCPorts.Transaction

	// Create Listen
	lis, err := net.Listen("tcp", transactionPort)
	if err != nil {
		log.Fatalf("err while create listen %v", err)
	}

	// Create Service Server
	s := grpc.NewServer()
	transactionpb.RegisterTransactionServiceServer(s, &Node{})

	log.Println(" gRPC server started on port:" + transactionPort)
	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("err while %v", err)
	}
}
