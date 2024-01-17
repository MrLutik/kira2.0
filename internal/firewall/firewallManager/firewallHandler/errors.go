package firewallHandler

import "errors"

var (
	ErrInvalidIPAddress = errors.New("invalid IP address")
	ErrInvalidPortType  = errors.New("port type is not valid")
	ErrNoPortMatch      = errors.New("no port matches found")
)
