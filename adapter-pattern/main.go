package main

import (
	"adapter-pattern/adapter"
	"adapter-pattern/computer"
	"fmt"
)

func main() {

	fmt.Println("RUN....")
	client := &Client{}

	mac := &computer.Mac{}
	client.InsertLightningConnectorIntoComputer(mac)

	windows := &computer.Windows{}
	windowsAdapter := &adapter.WindowsAdapter{WindowsMachine: windows}

	client.InsertLightningConnectorIntoComputer(windowsAdapter)

}
