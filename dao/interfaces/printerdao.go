package interfaces

import "github.com/qwertypomy/printers/models"

type PrinterDao interface {
	// PrintSize functions
	CreatePrintSize(printSize *models.PrintSize) (err error)
	UpdatePrintSize(printSize *models.PrintSize) (err error)
	DeletePrintSizeByID(id string) (err error)
	GetPrintSizeByID(id string) (printSize *models.PrintSize, err error)
	PrintSizeList() (printSizeList []models.PrintSize, err error)
	DeleteAllPrintSizes() (err error)
	// Brand functions
	CreateBrand(brand *models.Brand) (err error)
	UpdateBrand(brand *models.Brand) (err error)
	DeleteBrandByID(id string) (err error)
	GetBrandByID(id string) (brand *models.Brand, err error)
	BrandList() (brandList []models.Brand, err error)
	DeleteAllBrands() (err error)
	// PrintingTechnology functions
	CreatePrintingTechnology(printingTechnology *models.PrintingTechnology) (err error)
	UpdatePrintingTechnology(printingTechnology *models.PrintingTechnology) (err error)
	DeletePrintingTechnologyByID(id string) (err error)
	GetPrintingTechnologyByID(id string) (printingTechnology *models.PrintingTechnology, err error)
	PrintingTechnologyList() (printingTechnologyList []models.PrintingTechnology, err error)
	DeleteAllPrintingTechnologies() (err error)
	// FunctionType functions
	CreateFunctionType(functionType *models.FunctionType) (err error)
	UpdateFunctionType(functionType *models.FunctionType) (err error)
	DeleteFunctionTypeByID(id string) (err error)
	GetFunctionTypeByID(id string) (functionType *models.FunctionType, err error)
	FunctionTypeList() (functionTypeList []models.FunctionType, err error)
	DeleteAllFunctionTypes() (err error)
	// PrintResolution functions
	CreatePrintResolution(printResolution *models.PrintResolution) (err error)
	DeletePrintResolution(printResolution *models.PrintResolution) (err error)
	GetPrintResolution(printResolution *models.PrintResolution) (err error)
	PrintResolutionList() (printResolutionList []models.PrintResolution, err error)
	DeleteAllPrintResolutions() (err error)
	// Printer functions
	CreatePrinter(printer *models.Printer) (err error)
	UpdatePrinter(printer *models.Printer) (err error)
	DeletePrinterByID(id string) (err error)
	GetPrinterByID(id string) (printer *models.Printer, err error)
	PrinterList() (printerList []models.Printer, err error)
	DeleteAllPrinters() (err error)

	DeleteAllData() (err error)
}
