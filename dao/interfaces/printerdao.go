package interfaces

import "github.com/qwertypomy/printers/models"

type PrinterDao interface {
	CreatePrinterPrintSize(value string) (printSize models.PrinterPrintSize, err error)
	PrinterPrintSizeList() (printSizeList []models.PrinterPrintSize, err error)
	CreatePrinterConnectivityType(value string) (connectivityType models.PrinterConnectivityType, err error)
	PrinterConnectivityTypeList() (connectivityTypeList []models.PrinterConnectivityType, err error)
	CreatePrinterBrand(name string, description string) (brand models.PrinterBrand, err error)
	PrinterBrandList() (brandList []models.PrinterBrand, err error)
	CreatePrinterPrintingTechnology(name string, description string) (printingTechnology models.PrinterPrintingTechnology, err error)
	PrinterPrintingTechnologyList() (printingTechnologyList []models.PrinterPrintingTechnology, err error)
	CreatePrinterFunctionType(name string, description string) (functionType models.PrinterFunctionType, err error)
	PrinterFunctionTypeList() (functionTypeList []models.PrinterFunctionType, err error)
	CreatePrinterPrintResolution(x, y int) (printResolution models.PrinterPrintResolution, err error)
	PrinterPrintResolutionList() (printResolutionList []models.PrinterPrintResolution, err error)
}
