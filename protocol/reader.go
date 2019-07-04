package protocol

import (
	"bufio"
	"io"
	"log"
)

type CommandReader struct {
	reader *bufio.Reader
}

func NewCommandReader(reader io.Reader) *CommandReader {
	return &CommandReader{
		reader: bufio.NewReader(reader),
	}
}

func (r *CommandReader) Read() (interface{}, error) {
	// Read the first part
	commandName, err := r.reader.ReadString(' ')
	if err != nil {
		return nil, err
	}

	switch commandName {
	case "NAME ":
		user, err := r.reader.ReadString('\n')
		if err != nil {
			return nil, err
		}

		return NameCommand{
			Name: user[:len(user)-1],
		}, nil
	case "MESSAGE ":
		user, err := r.reader.ReadString(' ')
		if err != nil {
			return nil, err
		}

		message, err := r.reader.ReadString('\n')

		if err != nil {
			return nil, err
		}

		return MessageCommand{
			user[:len(user)-1],
			message[:len(message)-1],
		}, nil
	default:
		log.Printf("Unknown protocol: %v", commandName)
	}

	return nil, UnknownCommand{}
}
