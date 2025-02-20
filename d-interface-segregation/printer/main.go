package main

import (
	"fmt"
	"io"
	"os"
)

// Interface Segregation Principle (ISP) Example

// Problem: A "thick" interface (ThickPrinter),
// forces implementation of method that are not needed.

// .............................................................................
// ThickPrinter interface violates ISP (too broad)
type ThickPrinter interface {
	PrintDocument(document string)
	FaxDocument(document string)
	ScanDocument(document string)
}

// MultifunctionPrinter implements ThickPrinter
type MultifunctionPrinter struct{}

func (m MultifunctionPrinter) PrintDocument(document string) {
	fmt.Println("Printing:", document)
}

func (m MultifunctionPrinter) FaxDocument(document string) {
	fmt.Println("Faxing:", document)
}

func (m MultifunctionPrinter) ScanDocument(document string) {
	fmt.Println("Scanning:", document)
}

// SimplePrinter only needs to print,
// but must implement Fax and Scan (ISP violation)
type SimplePrinter struct{}

func (s SimplePrinter) PrintDocument(document string) {
	fmt.Println("Printing:", document)
}

func (s SimplePrinter) FaxDocument(document string) {
	// Forcing implementation is the problem
	panic("Faxing not supported")
}

func (s SimplePrinter) ScanDocument(document string) {
	// Forcing implementation is the problem
	panic("Scanning not supported")
}

// .............................................................................
// Solution: Segregate the interfaces into smaller, more specific ones.

// SPrinter interface (for printing only)
type SPrinter interface {
	PrintDocument(document string)
}

// SFax interface (for faxing only)
type SFax interface {
	FaxDocument(document string)
}

// SScanner interface (for scanning only)
type SScanner interface {
	ScanDocument(document string)
}

// MultifunctionPrinter (implements all interfaces)
type MultifunctionSPrinter struct{}

func (m MultifunctionSPrinter) PrintDocument(document string) {
	fmt.Println("Printing:", document)
}

func (m MultifunctionSPrinter) FaxDocument(document string) {
	fmt.Println("Faxing:", document)
}

func (m MultifunctionSPrinter) ScanDocument(document string) {
	fmt.Println("Scanning:", document)
}

// SimpleSPrinter (only implements SPrinter)
type SimpleSPrinter struct{}

func (s SimpleSPrinter) PrintDocument(document string) {
	fmt.Println("Printing:", document)
}

// .............................................................................
// Flexible and Composable solution with a standard interface

// APrinter interface (for printing, faxing, or scanning)
type APrinter interface {
	io.Writer
	Print()
}

type Printer struct {
	buf string
}

func (printer *Printer) Write(p []byte) (n int, err error) {
	// Append the byte slice to the string buffer
	printer.buf += string(p)
	// Return the number of bytes written (always successful here)
	return len(p), nil
}

func (printer *Printer) Print() {
	fmt.Printf("Printed: %s\n", printer.buf)
}

type Fax struct {
	buf string
}

func (printer *Fax) Write(p []byte) (n int, err error) {
	// Append the byte slice to the string buffer
	printer.buf += string(p)
	// Return the number of bytes written (always successful here)
	return len(p), nil
}

func (printer *Fax) Print() {
	fmt.Printf("Faxed: %s\n", printer.buf)
}

type AScanner interface {
	io.Reader
}

type Scanner struct {
	input  string
	offset int
}

func (s *Scanner) Read(p []byte) (n int, err error) {
	if s.offset >= len(s.input) {
		return 0, io.EOF // End of string
	}

	// Copy up to len(p) bytes or remaining string
	n = copy(p, s.input[s.offset:])

	s.offset += n
	return n, nil
}

func NewPipeline(scanner io.Reader, printers ...io.Writer) error {
	// Scan
	b, err := io.ReadAll(scanner)
	if err != nil {
		return err
	}
	fmt.Printf("Scanned: %s\n", string(b))

	// Write scanned input to printers
	for _, p := range printers {
		_, err = p.Write(b)
		if err != nil {
			return err
		}
	}

	// Print
	for _, p := range printers {
		if printer, ok := p.(APrinter); ok {
			printer.Print()
		}
	}

	return nil
}

// Example usage
func main() {
	scanner := &Scanner{input: "Hello World!"}

	printers := make([]io.Writer, 0)
	printers = append(printers, &Printer{})
	printers = append(printers, &Fax{})

	err := NewPipeline(scanner, printers...)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	os.Exit(0)
}
