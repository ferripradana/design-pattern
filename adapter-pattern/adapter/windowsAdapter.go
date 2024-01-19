package adapter

import (
	"adapter-pattern/computer"
	"fmt"
)

type WindowsAdapter struct {
	WindowsMachine *computer.Windows
}

func (w *WindowsAdapter) InsertIntoLightningPort() {
	fmt.Println("Adapter converts Lightning signal to USB.")
	w.WindowsMachine.InsertIntoUSBPort()
}
