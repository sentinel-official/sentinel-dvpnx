package utils

import (
	"crypto/rand"
	"crypto/tls"
	"fmt"
	"net"
	"net/http"

	"github.com/soheilhy/cmux"
)

// ListenAndServeTLS sets up a server that listens for both TLS and non-TLS traffic on the same address.
func ListenAndServeTLS(addr, certFile, keyFile string, handler http.Handler) error {
	// Load the TLS certificate and key
	cert, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		return fmt.Errorf("failed to load tls certificate: %w", err)
	}

	// Create a TCP listener
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return fmt.Errorf("failed to listen: %w", err)
	}

	// Create a cmux multiplexer
	mux := cmux.New(listener)

	// Define matchers for TLS and non-TLS traffic
	tlsMux := mux.Match(cmux.TLS())
	anyMux := mux.Match(cmux.Any())

	// Reuse the TLS configuration
	cfg := &tls.Config{
		Certificates: []tls.Certificate{cert},
		Rand:         rand.Reader,
	}

	// Serve TLS traffic
	go func() {
		tlsMux := tls.NewListener(tlsMux, cfg)
		if err := http.Serve(tlsMux, handler); err != nil {
			panic(fmt.Errorf("failed to serve tls: %w", err))
		}
	}()

	// Serve non-TLS traffic
	go func() {
		if err := http.Serve(anyMux, handler); err != nil {
			panic(fmt.Errorf("failed to serve any: %w", err))
		}
	}()

	// Start the multiplexer
	if err := mux.Serve(); err != nil {
		return fmt.Errorf("failed to serve: %w", err)
	}

	return nil
}
