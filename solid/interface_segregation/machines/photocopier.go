package machines

import "github.com/nachoberaza/GO_Desing_Patterns/solid/interface_segregation/document"

type Photocopier struct {
	printer Printer
	scanner Scanner
}

func (p *Photocopier) PrintDocument(d *document.Document) {
	p.printer.PrintNew(d)
}

func (p *Photocopier) ScanDocument(d *document.Document) {
	p.scanner.ScanNew(d)
}
