package main

import (
	"adapter-pattern/computer"
	"fmt"
)

type Client struct {
}

func (c *Client) InsertLightningConnectorIntoComputer(com computer.Computer) {
	fmt.Println("Client inserts Lightning connector into computer.")
	com.InsertIntoLightningPort()
}
