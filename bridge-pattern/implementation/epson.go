package implementation

import "fmt"

type Epson struct {
}

func (e *Epson) PrintFile() {
	fmt.Println("Printing by a Epson Printer")
}
