package abstraction

import (
	"bridge-pattern/implementation"
	"fmt"
)

type Windows struct {
	implementation.Printer
}

func (w *Windows) Print() {
	fmt.Println("Print request for windows")
	w.Printer.PrintFile()
}

func (w *Windows) SetPrinter(printer implementation.Printer) {
	w.Printer = printer
}
