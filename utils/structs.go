package utils

import {
	"fmt"
	"os"
	"net"
}

type MessageType string

const (
	Request MessageType = "request"
	Reply	MessageType = "reply"
)

type keepalive struct {
	SourceIP	net.IP
	DestIP		net.IP
	Type		MessageType
}

