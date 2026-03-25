package printer

import (
	"fmt"
	"io"
)

type Printer struct {
	Stdout  io.Writer
	Stderr  io.Writer
	Verbose bool
}

func NewPrinter(out, err io.Writer) *Printer {
	return &Printer{Stdout: out, Stderr: err}
}

func (p *Printer) Printf(format string, args ...any) {
	_, _ = fmt.Fprintf(p.Stdout, format, args...)
}

func (p *Printer) Println(args ...any) {
	_, _ = fmt.Fprintln(p.Stdout, args...)
}

func (p *Printer) Errorf(format string, args ...any) {
	_, _ = fmt.Fprintf(p.Stderr, format, args...)
}

func (p *Printer) Errorln(args ...any) {
	_, _ = fmt.Fprintln(p.Stderr, args...)
}

func (p *Printer) Verbosef(format string, args ...any) {
	if p.Verbose {
		_, _ = fmt.Fprintf(p.Stderr, format, args...)
	}
}

func (p *Printer) Verboseln(args ...any) {
	if p.Verbose {
		_, _ = fmt.Fprintln(p.Stderr, args...)
	}
}
