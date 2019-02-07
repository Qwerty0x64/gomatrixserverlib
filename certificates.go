package gomatrixserverlib

import (
	"crypto/x509"
	"errors"
	"net"
	"strings"
)

// IsValidCertificate checks if the given x509 certificate can be verified using
// system root CAs and an optional pool of intermediate CAs.
func IsValidCertificate(serverName ServerName, c *x509.Certificate, intermediates *x509.CertPool) (valid bool, err error) {
	// Fail if serverName is an ipv4/6
	if net.ParseIP(string(serverName)) != nil {
		// IP Address
		return false, errors.New("serverName is an IP literal")
	}

	// Remove port from serverName if it exists
	serverNameCleaned := strings.Split(string(serverName), ":")[0]

	// Check certificate chain validity
	verificationOpts := x509.VerifyOptions{
		DNSName:       serverNameCleaned,
		Intermediates: intermediates,
	}
	roots, err := c.Verify(verificationOpts)

	return len(roots) > 0, err
}