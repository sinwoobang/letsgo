package main

import (
	"bufio"
	"fmt"
	"letsgo/client"
	"log"
	"os"
	"strconv"
)

func main() {
	c := client.NewTcpChatClient()
	err := c.Dial(":3333")
	if err != nil {
		panic(err)
	}
	log.Print("Start Client")
	go c.Start()

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Enter Command(1. Set Name / 2. Send Message) : ")
		rawInput, _, _ := reader.ReadLine()
		commandNum, _ := strconv.Atoi(string(rawInput))
		if commandNum == 1 {
			fmt.Print("Name : ")
			rawInput, _, _ := reader.ReadLine()
			name, _ := strconv.Atoi(string(rawInput))
			c.SetName(string(name))
		} else {
			fmt.Print("Message : ")
			rawInput, _, _ := reader.ReadLine()
			name, _ := strconv.Atoi(string(rawInput))
			c.SendMessage(string(name))
		}
	}
}
