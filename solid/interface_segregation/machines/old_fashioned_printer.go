package machines

type OldFashionedPrinter struct {
}

func (o *OldFashionedPrinter) PrintOld() {
}

func (o *OldFashionedPrinter) FaxOld() {
	panic("operation not supported")
}

func (o *OldFashionedPrinter) ScanOld() {
	panic("operation not supported")
}
