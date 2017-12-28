package interfaces

import "github.com/qwertypomy/printers/models"

type PrinterDao interface {
	// PrintSize functions
	CreatePrintSize(printSize *models.PrintSize) (err error)
	UpdatePrintSize(printSize *models.PrintSize) (err error)
	DeletePrintSizeByID(id uint) (err error)
	GetPrintSizeByID(id uint) (printSize *models.PrintSize, err error)
	PrintSizeList() (printSizeList []models.PrintSize, err error)
	DeleteAllPrintSizes() (err error)
	// Brand functions
	CreateBrand(brand *models.Brand) (err error)
	UpdateBrand(brand *models.Brand) (err error)
	DeleteBrandByID(id uint) (err error)
	GetBrandByID(id uint) (brand *models.Brand, err error)
	BrandList() (brandList []models.Brand, err error)
	DeleteAllBrands() (err error)
	// PrintingTechnology functions
	CreatePrintingTechnology(printingTechnology *models.PrintingTechnology) (err error)
	UpdatePrintingTechnology(printingTechnology *models.PrintingTechnology) (err error)
	DeletePrintingTechnologyByID(id uint) (err error)
	GetPrintingTechnologyByID(id uint) (printingTechnology *models.PrintingTechnology, err error)
	PrintingTechnologyList() (printingTechnologyList []models.PrintingTechnology, err error)
	DeleteAllPrintingTechnologies() (err error)
	// FunctionType functions
	CreateFunctionType(functionType *models.FunctionType) (err error)
	UpdateFunctionType(functionType *models.FunctionType) (err error)
	DeleteFunctionTypeByID(id uint) (err error)
	GetFunctionTypeByID(id uint) (functionType *models.FunctionType, err error)
	FunctionTypeList() (functionTypeList []models.FunctionType, err error)
	DeleteAllFunctionTypes() (err error)
	// PrintResolution functions
	CreatePrintResolution(printResolution *models.PrintResolution) (err error)
	UpdatePrintResolution(printResolution *models.PrintResolution) (err error)
	DeletePrintResolution(printResolution *models.PrintResolution) (err error)
	GetPrintResolution(printResolution *models.PrintResolution) (err error)
	PrintResolutionList() (printResolutionList []models.PrintResolution, err error)
	DeleteAllPrintResolutions() (err error)
	// Printer functions
	CreatePrinter(printer *models.Printer) (err error)
	UpdatePrinter(printer *models.Printer) (err error)
	DeletePrinterByID(id uint) (err error)
	GetPrinterByID(id uint) (printer *models.Printer, err error)
	PrinterList() (printerList []models.Printer, err error)
	DeleteAllPrinters() (err error)

	DeleteAllData() (err error)
}
