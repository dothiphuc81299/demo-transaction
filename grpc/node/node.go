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

	// Get transactions by userID
	data, err := getTransactionDetailByUserID(req.GetUserID())

	result := &transactionpb.GetTransactionDetailByUserIDResponse{
		TransactionDetail: data,
	}
	return result, err
}

// GetTransactionDetailByCompanyID ...
func (s *Node) GetTransactionDetailByCompanyID(ctx context.Context, req *transactionpb.GetTransactionDetailByCompanyIDRequest) (*transactionpb.GetTransactionDetailByCompanyIDResponse, error) {

	// Get transaction by companyID
	data, err := getTransactionDetailByCompanyID(req.GetCompanyID())

	result := &transactionpb.GetTransactionDetailByCompanyIDResponse{
		TransactionDetail: data,
	}
	return result, err
}

// Start ...
func Start() {
	envVars := config.GetEnv()
	transactionPort := envVars.GRPCPorts.Transaction

	// Create listen
	lis, err := net.Listen("tcp", transactionPort)
	if err != nil {
		log.Fatalf("err while create listen %v", err)
	}

	// Create service server
	s := grpc.NewServer()
	transactionpb.RegisterTransactionServiceServer(s, &Node{})

	// Start server
	log.Println(" gRPC server started on port:" + transactionPort)
	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("err while %v", err)
	}
}
