package machines

import "github.com/nachoberaza/GO_Desing_Patterns/solid/interface_segregation/document"

type Machine interface {
	PrintOld(d *document.Document)
	FaxOld(d *document.Document)
	ScanOld(d *document.Document)
}

// These interface split the machines interface into two news
type Printer interface {
	PrintNew(d *document.Document)
}

type Scanner interface {
	ScanNew(d *document.Document)
}
