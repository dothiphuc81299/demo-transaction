package grpccompany

import (
	"log"

	"google.golang.org/grpc"

	"demo-transaction/config"
	companypb "demo-transaction/proto/models/company"
)

// CreateClient ...
func CreateClient() (*grpc.ClientConn, companypb.CompanyServiceClient) {
	envVars := config.GetEnv()
	address := envVars.GRPCAddresses.Company + envVars.GRPCPorts.Company

	// Create a client connection
	clientConn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("err while dial %v", err)
	}
	
	// Create company service 
	client := companypb.NewCompanyServiceClient(clientConn)
	return clientConn, client
}
