package abstraction

import "bridge-pattern/implementation"

type Computer interface {
	Print()
	SetPrinter(printer implementation.Printer)
}
