package machines

import "github.com/baldurdevs/go_desing_patterns/solid/interface_segregation/document"

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
