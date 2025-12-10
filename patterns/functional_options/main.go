package main

import (
	"fmt"
	"time"
)

// ============================================================================
// FUNCTIONAL OPTIONS PATTERN
// ============================================================================
// Patrón muy común en Go para configurar structs
// Más flexible que constructores con muchos parámetros
// Similar a Builder pattern en Java, pero más idiomático en Go
// ============================================================================

type Server struct {
	host    string
	port    int
	timeout time.Duration
	tls     bool
}

// Option es una función que modifica Server
type Option func(*Server)

// Funciones constructoras de opciones
func WithHost(host string) Option {
	return func(s *Server) {
		s.host = host
	}
}

func WithPort(port int) Option {
	return func(s *Server) {
		s.port = port
	}
}

func WithTimeout(timeout time.Duration) Option {
	return func(s *Server) {
		s.timeout = timeout
	}
}

func WithTLS(tls bool) Option {
	return func(s *Server) {
		s.tls = tls
	}
}

// NewServer crea un servidor con opciones
func NewServer(opts ...Option) *Server {
	// Valores por defecto
	s := &Server{
		host:    "localhost",
		port:    8080,
		timeout: 30 * time.Second,
		tls:     false,
	}

	// Aplicar opciones
	for _, opt := range opts {
		opt(s)
	}

	return s
}

func main() {
	// Uso básico con valores por defecto
	server1 := NewServer()
	fmt.Printf("Server 1: %+v\n", server1)

	// Con algunas opciones
	server2 := NewServer(
		WithHost("example.com"),
		WithPort(443),
		WithTLS(true),
	)
	fmt.Printf("Server 2: %+v\n", server2)

	// Con todas las opciones
	server3 := NewServer(
		WithHost("api.example.com"),
		WithPort(8080),
		WithTimeout(60*time.Second),
		WithTLS(false),
	)
	fmt.Printf("Server 3: %+v\n", server3)
}

