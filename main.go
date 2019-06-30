package main

import "letsgo/server"

func main() {
	s := server.NewTcpChatServer()
	s.Listen(":3333")
	// start the server
	s.Start()
}
