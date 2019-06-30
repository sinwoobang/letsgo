package protocol

type UnknownCommand struct {
}

func (u UnknownCommand) Error() string {
	return "Unknown Command"
}
