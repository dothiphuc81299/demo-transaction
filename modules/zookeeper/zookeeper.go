package zookeeper

import (
	"fmt"
	"os"
	"time"

	"github.com/samuel/go-zookeeper/zk"

	"demo-transaction/config"
)

var conn *zk.Conn

// Connect ...
func Connect() {
	var (
		uri     = os.Getenv("ZOOKEEPER_URI")
		envVars = config.GetEnv()
	)

	// Connect
	conn, _, err := zk.Connect([]string{uri}, time.Second*30)
	if err != nil {
		fmt.Println("ZookeeperURI:", uri)
		panic(err)
	}
	fmt.Println("Zookeeper Connected to", uri)

	// Get env key
	// App port
	appTransactionPort, _, _ := conn.Get("/app/port/transaction")
	envVars.AppPort = string(appTransactionPort)

	// Database
	databaseURI, _, _ := conn.Get("/database/uri")
	envVars.Database.URI = string(databaseURI)
	databaseTransactionName, _, _ := conn.Get("/database/name/transaction")
	envVars.Database.Name = string(databaseTransactionName)
	databaseTestName, _, _ := conn.Get("/database/test/transaction")
	envVars.Database.TestName = string(databaseTestName)

	// gRPCAddresses
	grpcAddressUser, _, _ := conn.Get("/grpc/uri/user")
	envVars.GRPCAddresses.User = string(grpcAddressUser)
	grpcAddressCompany, _, _ := conn.Get("/grpc/uri/company")
	envVars.GRPCAddresses.Company = string(grpcAddressCompany)
	grpcAddressTransaction, _, _ := conn.Get("/grpc/uri/transaction")
	envVars.GRPCAddresses.Transaction = string(grpcAddressTransaction)

	// gRPCPorts
	grpcPortUser, _, _ := conn.Get("/grpc/port/user")
	envVars.GRPCPorts.User = string(grpcPortUser)
	grpcPortCompany, _, _ := conn.Get("/grpc/port/company")
	envVars.GRPCPorts.Company = string(grpcPortCompany)
	grpcPortTransaction, _, _ := conn.Get("/grpc/port/transaction")
	envVars.GRPCPorts.Transaction = string(grpcPortTransaction)
}
