package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	conn, err := tls.DialWithDialer(
		&net.Dialer{Timeout: 30 * time.Second},
		"tcp",
		"www.youtube.com:443",
		&tls.Config{
			CurvePreferences: []tls.CurveID{tls.CurveP256},
			MinVersion:       tls.VersionTLS12,
		},
	)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	state := conn.ConnectionState()

	fmt.Printf("TLS Version: 1.%d\n", state.Version-tls.VersionTLS10)
	cipherSuiteName := tls.CipherSuiteName(state.CipherSuite)
	fmt.Println("Ciphersuite Name:", cipherSuiteName)
	issuerOrganization := state.VerifiedChains[0][0].Issuer.Organization[0]
	fmt.Println("Issuer Organization:", issuerOrganization)
}
