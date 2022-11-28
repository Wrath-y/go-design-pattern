package factory_simple

type Printer interface {
	Print(string) string
}

func NewPrinter(lang string) Printer {
	switch lang {
	case "cn":
		return new(CnPrinter)
	case "en":
		return new(EnPrinter)
	default:
		return new(CnPrinter)
	}
}

type CnPrinter struct{}

func (*CnPrinter) Print(name string) string {
	return name
}

type EnPrinter struct{}

func (*EnPrinter) Print(name string) string {
	return name
}
