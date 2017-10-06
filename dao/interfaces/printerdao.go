package interfaces

import "github.com/qwertypomy/printers/models"

type PrinterDao interface {
	Create(value interface{})
	Update(value interface{}, changedValue interface{})
	Delete(value interface{})

	GetPrintSizeById(id uint) (printSize *models.PrintSize)
	PrintSizeList() (printSizeList []models.PrintSize)
	DeleteAllPrintSizes()

	GetConnectivityTypeById(id uint) (connectivityType *models.ConnectivityType)
	ConnectivityTypeList() (connectivityTypeList []models.ConnectivityType)
	DeleteAllConnectivityTypes()

	GetBrandById(id uint) (brand *models.Brand)
	BrandList() (brandList []models.Brand)
	DeleteAllBrands()

	GetPrintingTechnologyById(id uint) (printingTechnology *models.PrintingTechnology)
	PrintingTechnologyList() (printingTechnologyList []models.PrintingTechnology)
	DeleteAllPrintingTechnologies()

	GetFunctionTypeById(id uint) (functionType *models.FunctionType)
	FunctionTypeList() (functionTypeList []models.FunctionType)
	DeleteAllFunctionTypes()

	GetPrintResolutionById(id uint) (printResolution *models.PrintResolution)
	PrintResolutionList() (printResolutionList []models.PrintResolution)
	DeleteAllPrintResolutions()

	GetPrinterById(id uint) (printer *models.Printer)
	PrinterList() (printerList []models.Printer)
	DeleteAllPrinters()

	DeleteAllData()
}
