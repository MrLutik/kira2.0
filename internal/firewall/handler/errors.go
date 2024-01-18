package handler

import "errors"

var (
	ErrInvalidIPAddress = errors.New("invalid IP address")
	ErrInvalidPortType  = errors.New("port type is not valid")
	ErrInvalidPort      = errors.New("invalid port")
)
