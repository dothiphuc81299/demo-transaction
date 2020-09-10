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
	var (
		userID = req.GetUserID()
	)

	// Get transaction by userID
	result, err := getTransactionDetailByUserID(userID)
	return result, err
}

// GetTransactionDetailByCompanyID ...
func (s *Node) GetTransactionDetailByCompanyID(ctx context.Context, req *transactionpb.GetTransactionDetailByCompanyIDRequest) (*transactionpb.GetTransactionDetailByCompanyIDResponse, error) {
	return nil, nil
}

// Start ...
func Start() {
	envVars := config.GetEnv()
	transactionPort := envVars.GRPCPorts.Transaction

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
