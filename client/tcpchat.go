package client

import (
	"io"
	"letsgo/protocol"
	"log"
	"net"
)

type TcpChatClient struct {
	conn      net.Conn
	cmdReader *protocol.CommandReader
	cmdWriter *protocol.CommandWriter
	name      string
	incoming  chan protocol.MessageCommand // Handle
}

func NewTcpChatClient() *TcpChatClient {
	return &TcpChatClient{
		incoming: make(chan protocol.MessageCommand),
	}
}

func (c *TcpChatClient) Dial(address string) error {
	// Connect to server and set new Reader and Writer to a Client Struct.
	conn, err := net.Dial("tcp", address)
	if err == nil {
		c.conn = conn
	}
	c.cmdReader = protocol.NewCommandReader(conn)
	c.cmdWriter = protocol.NewCommandWriter(conn)
	return err
}

func (c *TcpChatClient) Send(command interface{}) error {
	// Send a message which is based on Protocols. wrapped by the method SendMessage.
	return c.cmdWriter.Write(command)
}

func (c *TcpChatClient) Start() {
	// Start a communication after Dial(). Handle incoming messages.
	for {
		cmd, err := c.cmdReader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Printf("Read error %v", err)
		}
		if cmd != nil {
			switch v := cmd.(type) {
			case protocol.MessageCommand:
				c.incoming <- v
			default:
				log.Printf("Unknown command: %v", v)
			}
		}
	}
}
