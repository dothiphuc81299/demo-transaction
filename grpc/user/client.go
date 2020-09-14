package grpcuser

import (
	"log"

	"google.golang.org/grpc"

	"demo-transaction/config"
	userpb "demo-transaction/proto/models/user"
)

// CreateClient ...
func CreateClient() (*grpc.ClientConn, userpb.UserServiceClient) {
	envVars := config.GetEnv()
	address := envVars.GRPCAddresses.User + envVars.GRPCPorts.User

	// Create a client connection
	clientConn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("err while dial %v", err)
	}

	// Create user service 
	client := userpb.NewUserServiceClient(clientConn)
	return clientConn, client
}
