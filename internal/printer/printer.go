package printer

import "io"

type Printer struct {
	Out io.Writer
}

func NewPrinter(out io.Writer) *Printer {
	return &Printer{Out: out}
}
