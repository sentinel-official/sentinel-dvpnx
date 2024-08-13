package utils

import (
	"crypto/rand"
	"crypto/tls"
	"net"
	"net/http"

	"github.com/soheilhy/cmux"
)

// ListenAndServeTLS sets up a server that listens for both TLS and non-TLS traffic on the same address.
func ListenAndServeTLS(addr, certFile, keyFile string, handler http.Handler) error {
	// Load the TLS certificate and key
	cert, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		return err
	}

	// Create a TCP listener
	l, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	// Create a cmux multiplexer
	mux := cmux.New(l)

	// Create matchers for TLS and non-TLS traffic
	tlsMux := mux.Match(cmux.TLS())
	anyMux := mux.Match(cmux.Any())

	// Serve TLS traffic
	go func() {
		err := http.Serve(
			tls.NewListener(tlsMux, &tls.Config{
				Certificates: []tls.Certificate{cert},
				Rand:         rand.Reader,
			}),
			handler,
		)
		if err != nil {
			panic(err)
		}
	}()

	// Serve non-TLS traffic
	go func() {
		err := http.Serve(anyMux, handler)
		if err != nil {
			panic(err)
		}
	}()

	// Start the mux server
	return mux.Serve()
}
