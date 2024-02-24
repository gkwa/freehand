package functional_options2

import (
	"testing"
)

func TestNewServerConfig_DefaultConfig(t *testing.T) {
	config := NewServerConfig()

	if config.Host != "localhost" {
		t.Errorf("Expected default host to be localhost, got %s", config.Host)
	}

	if config.Port != 8080 {
		t.Errorf("Expected default port to be 8080, got %d", config.Port)
	}

	if config.Protocol != "http" {
		t.Errorf("Expected default protocol to be http, got %s", config.Protocol)
	}
}

func TestNewServerConfig_WithOptions(t *testing.T) {
	config := NewServerConfig(
		WithHost("example.com"),
		WithPort(9000),
		WithProtocol("https"),
	)

	if config.Host != "example.com" {
		t.Errorf("Expected host to be example.com, got %s", config.Host)
	}

	if config.Port != 9000 {
		t.Errorf("Expected port to be 9000, got %d", config.Port)
	}

	if config.Protocol != "https" {
		t.Errorf("Expected protocol to be https, got %s", config.Protocol)
	}
}

func TestNewServer_DefaultConfig(t *testing.T) {
	server := NewServer()

	if server.Config.Host != "localhost" {
		t.Errorf("Expected default host to be localhost, got %s", server.Config.Host)
	}

	if server.Config.Port != 8080 {
		t.Errorf("Expected default port to be 8080, got %d", server.Config.Port)
	}

	if server.Config.Protocol != "http" {
		t.Errorf("Expected default protocol to be http, got %s", server.Config.Protocol)
	}
}

func TestNewServer_WithOptions(t *testing.T) {
	server := NewServer(
		WithHost("example.com"),
		WithPort(9000),
		WithProtocol("https"),
	)

	if server.Config.Host != "example.com" {
		t.Errorf("Expected host to be example.com, got %s", server.Config.Host)
	}

	if server.Config.Port != 9000 {
		t.Errorf("Expected port to be 9000, got %d", server.Config.Port)
	}

	if server.Config.Protocol != "https" {
		t.Errorf("Expected protocol to be https, got %s", server.Config.Protocol)
	}
}
