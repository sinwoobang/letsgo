package server_test

import (
	"fmt"
	"letsgo/server"
	"testing"
)

func TestNewTcpChatServer(t *testing.T) {
	s := server.NewTcpChatServer()
	if s == nil {
		t.Error("TcpChatServer is nil")
	}
}

func TestTcpChatServer_Listen(t *testing.T) {
	s := server.NewTcpChatServer()
	err := s.Listen("127.0.0.1:3030")
	if err != nil {
		t.Error(fmt.Sprintf("Error in Listen() : %s", err))
	}
}

func TestTcpChatServer_Close(t *testing.T) {
	// Since Close() has no return value, it calls Start() to validate whether it is closed.
	defer func() {
		if err := recover(); err == nil {
			t.Error("Close() doesn't work.")
		}
	}()

	s := server.NewTcpChatServer()
	_ = s.Listen("127.0.0.1:3030")
	s.Close()

	s.Start() // Need to be panic in a normal case cause it's called after Close().
}
