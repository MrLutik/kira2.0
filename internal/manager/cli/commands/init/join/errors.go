package join

import "errors"

var (
	ErrInvalidSekaidP2PPort = errors.New("invalid Sekaid P2P port")
	ErrInvalidSekaidRPCPort = errors.New("invalid Sekaid RPC port")
	ErrInvalidInterxPort    = errors.New("invalid Interx port")
	ErrInvalidIPAddress     = errors.New("invalid IP address")
)
