package models

// Abstract struct IdValue
type IdValue struct {
	Id    int
	Value string
}

// Abstract struct IdNameDescription
type IdNameDescription struct {
	Id          int
	Name        string
	Description string
}

type Printer struct {
	Id                 int
	Name               string
	Description        string
	PagePerMinute      int
	Brand              *PrinterBrand
	PrintingTechnology *PrinterPrintingTechnology
	FunctionType       *PrinterFunctionType
	PrintSize          *PrinterPrintSize
	PrintResolution    *PrinterPrintResolution
	ConnectivityType   []PrinterConnectivityType
	Size               string
	Weight             string
	AdditionalInfo     string
	Number             int
	Price              int
}

type PrinterBrand struct {
	IdNameDescription
}

type PrinterPrintingTechnology struct {
	IdNameDescription
}

type PrinterFunctionType struct {
	IdNameDescription
}

type PrinterPrintSize struct {
	IdValue
}

type PrinterPrintResolution struct {
	X int
	Y int
}

type PrinterConnectivityType struct {
	IdValue
}
