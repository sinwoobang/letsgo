package client

import (
	"letsgo/protocol"
)

type ChatClient interface {
	Dial(address string) error // Connect to server
	Send(command interface{}) error
	SendMessage(message string) error
	SetName(name string) error // Set display name
	Start()
	Close()
	Incoming() chan protocol.MessageCommand // Handle messages from server
}
