package controller

import (
	"errors"

	"github.com/mrlutik/kira2.0/internal/types"
)

var (
	ErrEmptyZone       = errors.New("zone name cannot be empty")
	ErrEmptyIPAddress  = errors.New("IP address cannot be empty")
	ErrEmptyPortNumber = errors.New("port number cannot be empty")
	ErrInvalidPortType = errors.New("port type must be 'tcp' or 'udp'")
	ErrEmptyInterface  = errors.New("interface name cannot be empty")
)

// validateZoneName checks if the provided zone name is valid.
func validateZoneName(zoneName string) error {
	if zoneName == "" {
		return ErrEmptyZone
	}
	return nil
}

// validatePort checks if the provided port is valid.
func validatePort(port types.Port) error {
	if port.Port == "" {
		return ErrEmptyPortNumber
	}
	if port.Type != "tcp" && port.Type != "udp" {
		return ErrInvalidPortType
	}
	return nil
}

// validateIP checks if the provided IP address is valid.
func validateIP(ip string) error {
	if ip == "" {
		return ErrEmptyIPAddress
	}
	return nil
}

// validateInterfaceName checks if the provided interface name is valid.
func validateInterfaceName(interfaceName string) error {
	if interfaceName == "" {
		return ErrEmptyInterface
	}
	return nil
}
