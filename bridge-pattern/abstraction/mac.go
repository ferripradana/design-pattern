package abstraction

import (
	"bridge-pattern/implementation"
	"fmt"
)

type Mac struct {
	implementation.Printer
}

func (m *Mac) Print() {
	fmt.Println("Print request for mac")
	m.Printer.PrintFile()
}

func (m *Mac) SetPrinter(printer implementation.Printer) {
	m.Printer = printer
}
