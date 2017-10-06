package mysql

import (
	"github.com/qwertypomy/printers/models"
)

type PrinterDaoImpl struct {
}

// General Create, Update, Delete function for all entities
func (PrinterDaoImpl) Create(value interface{}) {
	Db.Create(value)
}

func (PrinterDaoImpl) Update(value interface{}, changedValue interface{}) {
	Db.Model(value).Omit("id").Updates(changedValue)
}

func (PrinterDaoImpl) Delete(value interface{}) {
	Db.Delete(value)
}

// PrintSize functions
func (PrinterDaoImpl) GetPrintSizeById(id uint) (printSize *models.PrintSize) {
	Db.Where("id = ?", id).First(&printSize)
	return
}

func (PrinterDaoImpl) PrintSizeList() (printSizeList []models.PrintSize) {
	Db.Find(&printSizeList)
	return
}

func (PrinterDaoImpl) DeleteAllPrintSizes() {
	Db.Exec("truncate table print_sizes")
}

// ConnectivityType functions
func (PrinterDaoImpl) GetConnectivityTypeById(id uint) (connectivityType *models.ConnectivityType) {
	Db.Where("id = ?", id).First(&connectivityType)
	return
}

func (PrinterDaoImpl) ConnectivityTypeList() (connectivityTypeList []models.ConnectivityType) {
	Db.Find(&connectivityTypeList)
	return
}

func (PrinterDaoImpl) DeleteAllConnectivityTypes() {
	Db.Exec("truncate table connectivity_types")
}

// Brand functions
func (PrinterDaoImpl) GetBrandById(id uint) (brand *models.Brand) {
	Db.Where("id = ?", id).First(&brand)
	return
}

func (PrinterDaoImpl) BrandList() (brandList []models.Brand) {
	Db.Find(&brandList)
	return
}

func (PrinterDaoImpl) DeleteAllBrands() {
	Db.Exec("truncate table brands")
}

// PrintingTechnology functions
func (PrinterDaoImpl) GetPrintingTechnologyById(id uint) (printingTechnology *models.PrintingTechnology) {
	Db.Where("id = ?", id).First(&printingTechnology)
	return
}

func (PrinterDaoImpl) PrintingTechnologyList() (printingTechnologyList []models.PrintingTechnology) {
	Db.Find(&printingTechnologyList)
	return
}

func (PrinterDaoImpl) DeleteAllPrintingTechnologies() {
	Db.Exec("truncate table printing_technologies")
}

// FunctionType functions
func (PrinterDaoImpl) GetFunctionTypeById(id uint) (functionType *models.FunctionType) {
	Db.Where("id = ?", id).First(&functionType)
	return
}

func (PrinterDaoImpl) FunctionTypeList() (functionTypeList []models.FunctionType) {
	Db.Find(&functionTypeList)
	return
}

func (PrinterDaoImpl) DeleteAllFunctionTypes() {
	Db.Exec("truncate table function_types")
}

// PrintResolution functions
func (PrinterDaoImpl) GetPrintResolutionById(id uint) (printResolution *models.PrintResolution) {
	Db.Where("id = ?", id).First(&printResolution)
	return
}

func (PrinterDaoImpl) PrintResolutionList() (printResolutionList []models.PrintResolution) {
	Db.Find(&printResolutionList)
	return
}

func (PrinterDaoImpl) DeleteAllPrintResolutions() {
	Db.Exec("truncate table print_resolutions")
}

// Printer functions
func (PrinterDaoImpl) GetPrinterById(id uint) (printer *models.Printer) {
	Db.Where("id = ?", id).First(&printer)
	return
}

func (PrinterDaoImpl) PrinterList() (printerList []models.Printer) {
	Db.Find(&printerList)
	return
}

func (PrinterDaoImpl) DeleteAllPrinters() {
	Db.Exec("truncate table printers")
}

// All Tables functions
func (pdi PrinterDaoImpl) DeleteAllData() {
	pdi.DeleteAllPrintSizes()
	pdi.DeleteAllConnectivityTypes()
	pdi.DeleteAllBrands()
	pdi.DeleteAllPrintingTechnologies()
	pdi.DeleteAllFunctionTypes()
	pdi.DeleteAllPrintResolutions()
	pdi.DeleteAllPrinters()
}
