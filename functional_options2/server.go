package functional_options2

import (
	"fmt"
)

// ServerConfig contains configurations for a server.
type ServerConfig struct {
	Host     string
	Port     int
	Protocol string
}

// ServerOption defines a functional option for configuring a server.
type ServerOption func(*ServerConfig)

// WithHost sets the host for the server.
func WithHost(host string) ServerOption {
	return func(c *ServerConfig) {
		c.Host = host
	}
}

// WithPort sets the port for the server.
func WithPort(port int) ServerOption {
	return func(c *ServerConfig) {
		c.Port = port
	}
}

// WithProtocol sets the protocol for the server.
func WithProtocol(protocol string) ServerOption {
	return func(c *ServerConfig) {
		c.Protocol = protocol
	}
}

// NewServerConfig creates a new ServerConfig with the provided options.
func NewServerConfig(opts ...ServerOption) *ServerConfig {
	config := &ServerConfig{
		Host:     "localhost",
		Port:     8080,
		Protocol: "http",
	}

	for _, opt := range opts {
		opt(config)
	}

	return config
}

// Server represents a server.
type Server struct {
	Config *ServerConfig
	// Add other fields here...
}

// NewServer creates a new Server with the provided configurations.
func NewServer(opts ...ServerOption) *Server {
	return &Server{
		Config: NewServerConfig(opts...),
		// Initialize other fields here...
	}
}

func RunTest() {
	// Create a server with default configuration
	server1 := NewServer()
	fmt.Println("Server 1 Config:", server1.Config)

	// Create a server with custom configuration
	server2 := NewServer(
		WithHost("example.com"),
		WithPort(9000),
		WithProtocol("https"),
	)
	fmt.Println("Server 2 Config:", server2.Config)
}
