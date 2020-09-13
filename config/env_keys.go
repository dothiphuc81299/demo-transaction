package config

// ENV ...
type ENV struct {
	IsDev bool

	// Zookeeper URI
	ZookeeperURI string

	// App port
	AppPort string

	// Database
	Database struct {
		URI      string
		TestName string
		Name     string
	}

	// Redis
	Redis struct {
		URI  string
		Pass string
	}

	// gRPC addresses
	GRPCAddresses struct {
		User        string
		Transaction string
		Company     string
	}

	// gRPC ports
	GRPCPorts struct {
		User        string
		Transaction string
		Company     string
	}
}

var env ENV

// InitENV ...
func InitENV() {
	env = ENV{
		IsDev: true,
	}
}

// GetEnv ...
func GetEnv() *ENV {
	return &env
}
